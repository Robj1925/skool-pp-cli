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

			// We use the chat-channels endpoint which contains group info for each channel
			path := "/self/chat-channels"
			params := map[string]string{
				"limit": "100",
			}

			data, err := c.Get(path, params)
			if err != nil {
				return classifyAPIError(err, flags)
			}

			var channels []map[string]any
			if err := json.Unmarshal(data, &channels); err != nil {
				return fmt.Errorf("parsing response: %w", err)
			}

			type GroupInfo struct {
				Name string `json:"name"`
				Hash string `json:"hash"`
				Slug string `json:"slug"`
			}
			
			groupsMap := make(map[string]GroupInfo)
			for _, ch := range channels {
				if gAny, ok := ch["group"]; ok {
					if g, ok := gAny.(map[string]any); ok {
						name, _ := g["name"].(string)
						hash, _ := g["id"].(string)
						slug, _ := g["slug"].(string)
						if hash != "" {
							groupsMap[hash] = GroupInfo{Name: name, Hash: hash, Slug: slug}
						}
					}
				}
			}

			var groups []map[string]any
			for _, g := range groupsMap {
				groups = append(groups, map[string]any{
					"Name": g.Name,
					"Hash": g.Hash,
					"Slug": g.Slug,
				})
			}

			if flags.asJSON {
				return printJSONFiltered(cmd.OutOrStdout(), groups, flags)
			}

			if len(groups) == 0 {
				fmt.Fprintln(os.Stderr, "No groups found.")
				return nil
			}

			return printAutoTable(cmd.OutOrStdout(), groups)
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
			data, err := c.Get(path, nil)
			if err != nil {
				return classifyAPIError(err, flags)
			}

			return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
		},
	}
}
