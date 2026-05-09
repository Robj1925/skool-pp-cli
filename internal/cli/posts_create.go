// Copyright 2026 robj1925. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func newPostsCreateCmd(flags *rootFlags) *cobra.Command {
	var title, content, categoryID string

	cmd := &cobra.Command{
		Use:     "create <group-id>",
		Short:   "Create a new post in a group",
		Example: "  skool-pp-cli posts create 7e41737e8f404893aa7138fb01bc63e0 --title \"Hello\" --content \"World\"",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/posts"

			// Default to "General discussion" category if none provided
			// In a real app, we'd fetch this from the group's discovery data
			activeCategoryID := categoryID
			if activeCategoryID == "" {
				activeCategoryID = "fff39a66857540d6a6488c921f50f208"
			}

			body := map[string]any{
				"group_id":  args[0],
				"post_type": "generic",
				"metadata": map[string]any{
					"title":   title,
					"content": content,
					"action":  0,
					"labels":  activeCategoryID,
				},
			}

			data, statusCode, err := c.Post(path, body)
			if err != nil {
				return classifyAPIError(err, flags)
			}

			// Parse response to get the post ID and name
			var resp struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			}
			_ = json.Unmarshal(data, &resp)

			if flags.asJSON || !isTerminal(cmd.OutOrStdout()) {
				envelope := map[string]any{
					"action":   "post",
					"resource": "posts",
					"path":     path,
					"status":   statusCode,
					"success":  statusCode >= 200 && statusCode < 300,
				}
				if flags.dryRun {
					envelope["dry_run"] = true
					envelope["status"] = 0
					envelope["success"] = false
				} else {
					var respData any
					if err := json.Unmarshal(data, &respData); err == nil {
						envelope["data"] = respData
					}
				}
				return flags.printJSON(cmd, envelope)
			}

			if flags.dryRun {
				fmt.Println("Dry run: post creation request prepared.")
				return nil
			}

			fmt.Printf("Successfully created post!\n")
			fmt.Printf("  ID:   %s\n", resp.ID)
			fmt.Printf("  Slug: %s\n", resp.Name)
			if categoryID == "" {
				fmt.Printf("  Note: Defaulted to 'General discussion' category (%s)\n", activeCategoryID)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&title, "title", "", "Title of the post")
	cmd.Flags().StringVar(&content, "content", "", "Content of the post")
	cmd.Flags().StringVar(&categoryID, "category-id", "", "ID of the category (label)")

	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("content")

	return cmd
}
