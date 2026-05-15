package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newPostsListCategoriesCmd(flags *rootFlags) *cobra.Command {
	var groupHash string

	cmd := &cobra.Command{
		Use:   "list-categories",
		Short: "List post categories for a community",
		RunE: func(cmd *cobra.Command, args []string) error {
			if groupHash == "" {
				return fmt.Errorf("group hash is required (use --group)")
			}

			c, err := flags.newClient()
			if err != nil {
				return err
			}

			// We use the discovery endpoint which contains category info
			path := fmt.Sprintf("/groups/%s/discovery", groupHash)
			data, err := c.GetWithHeaders(path, nil, skoolHeaders())
			if err != nil {
				return classifyAPIError(err, flags)
			}

			var discovery map[string]any
			if err := json.Unmarshal(data, &discovery); err != nil {
				return fmt.Errorf("parsing response: %w", err)
			}

			// Extract categories from discovery data
			var categories []any
			if cats, ok := discovery["categories"].([]any); ok {
				categories = cats
			} else if group, ok := discovery["group"].(map[string]any); ok {
				if cats, ok := group["categories"].([]any); ok {
					categories = cats
				}
			}

			var catsFormatted []map[string]any
			for _, c := range categories {
				if m, ok := c.(map[string]any); ok {
					catsFormatted = append(catsFormatted, m)
				}
			}

			if flags.asJSON {
				return printJSONFiltered(cmd.OutOrStdout(), catsFormatted, flags)
			}

			if len(catsFormatted) == 0 {
				fmt.Fprintln(os.Stderr, "No categories found for this group.")
				return nil
			}

			return printAutoTable(cmd.OutOrStdout(), catsFormatted)
		},
	}

	cmd.Flags().StringVar(&groupHash, "group", "", "Group hash (ID)")
	return cmd
}
