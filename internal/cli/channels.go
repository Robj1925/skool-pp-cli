// Copyright 2026 robj1925. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newChannelsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "channels",
		Short: "Operations on chat channels",
	}

	cmd.AddCommand(newChannelsCreateMessageCmd(flags))
	cmd.AddCommand(newChannelsListCmd(flags))
	return cmd
}
