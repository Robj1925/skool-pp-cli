// Copyright 2026 robj1925. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newPostsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "posts",
		Short: "Manage community posts",
		Long:  `Create and manage posts in your Skool communities.`,
		Example: `  skool-pp-cli posts create <group-id> --title "My Post" --content "Body"`,
	}

	cmd.AddCommand(newPostsCreateCmd(flags))
	return cmd
}
