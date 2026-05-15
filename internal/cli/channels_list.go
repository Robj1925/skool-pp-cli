// Copyright 2026 robj1925. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

func newChannelsListCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List chat conversations",
		Example: "  skool-pp-cli channels list",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			// Correct path discovered from browser traffic
			path := "/self/chat-channels"
			
			params := map[string]string{
				"offset":      "0",
				"limit":       "30",
				"last":        "true",
				"unread-only": "false",
			}

			if dryRunOK(flags) {
				fmt.Printf("GET %s%s\n", c.BaseURL, path)
				return nil
			}

			headers := map[string]string{
				"Host":              "api2.skool.com",
				"Origin":            "https://www.skool.com",
				"Referer":           "https://www.skool.com/",
				"User-Agent":        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36",
				"Accept":            "*/*",
				"Accept-Language":   "en-US,en;q=0.9",
				"Sec-Fetch-Dest":    "empty",
				"Sec-Fetch-Mode":    "cors",
				"Sec-Fetch-Site":    "same-site",
				"Content-Type":      "application/json",
			}



			data, err := c.GetWithHeaders(path, params, headers)
			if err != nil {
				return classifyAPIError(err, flags)
			}

			return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
		},
	}

	return cmd
}
