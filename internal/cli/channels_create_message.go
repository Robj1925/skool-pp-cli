// Copyright 2026 robj1925. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

func newChannelsCreateMessageCmd(flags *rootFlags) *cobra.Command {
	var flagContent string
	var flagCt string

	cmd := &cobra.Command{
		Use:   "send <channel-id>",
		Short: "Send a message to a chat channel",
		Example: "  skool-pp-cli channels send 50513db7842047268f62dcd2b776533f --content \"Hello from CLI!\"",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			channelID := args[0]
			path := fmt.Sprintf("/channels/%s/messages", channelID)
			
			fullPath := path
			if flagCt != "" {
				fullPath = fmt.Sprintf("%s?ct=%s", path, flagCt)
			} else {
				fullPath = fmt.Sprintf("%s?ct=wdm", path)
			}

			body := map[string]any{
				"content":     flagContent,
				"attachments": []any{},
			}

			if dryRunOK(flags) {
				fmt.Printf("POST %s%s\n", c.BaseURL, fullPath)
				fmt.Printf("  Body: %v\n", body)
				return nil
			}

			headers := map[string]string{
				"Host":              "api2.skool.com",
				"Origin":            "https://www.skool.com",
				"Referer":           fmt.Sprintf("https://www.skool.com/chat?ch=%s", channelID),
				"User-Agent":        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
				"X-Requested-With":  "XMLHttpRequest",
				"Accept":            "application/json, text/plain, */*",
				"Accept-Language":   "en-US,en;q=0.9",
				"Sec-Ch-Ua":         `"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"`,
				"Sec-Ch-Ua-Mobile":  "?0",
				"Sec-Ch-Ua-Platform": `"macOS"`,
				"Sec-Fetch-Dest":    "empty",
				"Sec-Fetch-Mode":    "cors",
				"Sec-Fetch-Site":    "same-site",
			}

			data, _, err := c.PostWithHeaders(fullPath, body, headers)
			if err != nil {
				return classifyAPIError(err, flags)
			}

			return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
		},
	}

	cmd.Flags().StringVar(&flagContent, "content", "", "Message content")
	cmd.Flags().StringVar(&flagCt, "ct", "wdm", "Client type marker")
	_ = cmd.MarkFlagRequired("content")

	return cmd
}
