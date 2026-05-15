// Copyright 2026 robj1925. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newChannelsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "channels",
		Short: "Direct messaging and chat channels",
		Long:  `List chat conversations and send direct messages to community members.`,
		Example: `  skool-pp-cli channels list
  skool-pp-cli channels send <channel-id> --content "Hello"`,
	}

	cmd.AddCommand(newChannelsCreateMessageCmd(flags))
	cmd.AddCommand(newChannelsListCmd(flags))
	return cmd
}
