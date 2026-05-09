// Copyright 2026 robj1925. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newPostsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "posts",
		Short: "Operations on posts",
	}

	cmd.AddCommand(newPostsCreateCmd(flags))
	return cmd
}
