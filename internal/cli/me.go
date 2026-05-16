package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newMeCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "me",
		Short: "View your current Skool profile and groups",
	}

	cmd.AddCommand(newMeGroupsCmd(flags))
	cmd.AddCommand(newMeInfoCmd(flags))

	return cmd
}

func newMeGroupsCmd(flags *rootFlags) *cobra.Command {
	return &cobra.Command{
		Use:   "groups",
		Short: "List communities you belong to",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			// Endpoint Swap: Change from /self/chat-channels to /self/groups
			path := "/self/groups"
			data, err := c.GetWithHeaders(path, nil, skoolHeaders())
			if err != nil {
				return classifyAPIError(err, flags)
			}

			var groupsData []map[string]any
			if err := json.Unmarshal(data, &groupsData); err != nil {
				// Handle wrapped response { "groups": [...] }
				var wrapped struct {
					Groups []map[string]any `json:"groups"`
				}
				if err2 := json.Unmarshal(data, &wrapped); err2 == nil {
					groupsData = wrapped.Groups
				} else {
					return fmt.Errorf("parsing response: %w", err)
				}
			}

			var results []map[string]any
			for _, g := range groupsData {
				// Metadata Mapping: Use display_name from metadata or top-level
				name := ""
				if metadata, ok := g["metadata"].(map[string]any); ok {
					name, _ = metadata["display_name"].(string)
				}
				if name == "" {
					name, _ = g["display_name"].(string)
				}
				if name == "" {
					name, _ = g["name"].(string)
				}

				hash, _ := g["id"].(string)
				slug, _ := g["name"].(string)

				if hash != "" {
					results = append(results, map[string]any{
						"Name": name,
						"Hash": hash,
						"Slug": slug,
					})
				}
			}

			if flags.asJSON {
				return printJSONFiltered(cmd.OutOrStdout(), results, flags)
			}

			if len(results) == 0 {
				fmt.Fprintln(os.Stderr, "No groups found.")
				return nil
			}

			return printAutoTable(cmd.OutOrStdout(), results)
		},
	}
}

func newMeInfoCmd(flags *rootFlags) *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Show your Skool user profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/self"
			data, err := c.GetWithHeaders(path, nil, skoolHeaders())
			if err != nil {
				return classifyAPIError(err, flags)
			}

			return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
		},
	}
}
