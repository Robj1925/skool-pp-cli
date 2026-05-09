package cli

import (
	"encoding/json"
	"fmt"
	"github.com/Robj1925/skool-pp-cli/internal/store"
	"time"

	"github.com/spf13/cobra"
)

func newAnalyticsDomainCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "analytics-domain",
		Aliases: []string{"domain-analytics"},
		Short:   "Domain analytics commands (Rung 4)",
	}

	cmd.AddCommand(newStaleCmd(flags))
	cmd.AddCommand(newOrphansCmd(flags))
	cmd.AddCommand(newTopContributorsCmd(flags))
	return cmd
}

func newStaleCmd(flags *rootFlags) *cobra.Command {
	var days int
	cmd := &cobra.Command{
		Use:   "stale",
		Short: "Members with no recent activity",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := store.OpenReadOnly(flags.configPath + "/../skool.db")
			if err != nil {
				return fmt.Errorf("local store not available. Run 'skool-pp-cli sync' first. Error: %w", err)
			}
			defer db.Close()

			threshold := time.Now().AddDate(0, 0, -days).Format(time.RFC3339)
			rows, err := db.Query(`
				SELECT id, name, last_active_at, email
				FROM members 
				WHERE last_active_at < ?
				ORDER BY last_active_at DESC
			`, threshold)
			if err != nil {
				return err
			}
			defer rows.Close()

			var results []map[string]any
			for rows.Next() {
				var id, name, lastActive, email string
				if err := rows.Scan(&id, &name, &lastActive, &email); err == nil {
					results = append(results, map[string]any{
						"id":             id,
						"name":           name,
						"email":          email,
						"last_active_at": lastActive,
					})
				}
			}

			if flags.asJSON {
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(results)
			}

			fmt.Fprintln(cmd.OutOrStdout(), "MEMBER ID\tNAME\tEMAIL\tLAST ACTIVE")
			for _, r := range results {
				fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\t%s\n", r["id"], r["name"], r["email"], r["last_active_at"])
			}
			return nil
		},
	}
	cmd.Flags().IntVar(&days, "days", 30, "Number of days of inactivity")
	return cmd
}

func newOrphansCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "orphans",
		Short: "Members who joined but never posted",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := store.OpenReadOnly(flags.configPath + "/../skool.db")
			if err != nil {
				return fmt.Errorf("local store not available. Run 'skool-pp-cli sync' first. Error: %w", err)
			}
			defer db.Close()

			rows, err := db.Query(`
				SELECT id, name, joined_at, level
				FROM members 
				WHERE post_count = 0
				ORDER BY joined_at ASC
			`)
			if err != nil {
				return err
			}
			defer rows.Close()

			var results []map[string]any
			for rows.Next() {
				var id, name, joinedAt string
				var level int
				if err := rows.Scan(&id, &name, &joinedAt, &level); err == nil {
					results = append(results, map[string]any{
						"id":        id,
						"name":      name,
						"joined_at": joinedAt,
						"level":     level,
					})
				}
			}

			if flags.asJSON {
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(results)
			}

			fmt.Fprintln(cmd.OutOrStdout(), "MEMBER ID\tNAME\tLEVEL\tJOINED AT")
			for _, r := range results {
				fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%d\t%s\n", r["id"], r["name"], r["level"], r["joined_at"])
			}
			return nil
		},
	}
	return cmd
}

func newTopContributorsCmd(flags *rootFlags) *cobra.Command {
	var period string
	cmd := &cobra.Command{
		Use:   "top-contributors",
		Short: "Most active members based on points/levels",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := store.OpenReadOnly(flags.configPath + "/../skool.db")
			if err != nil {
				return fmt.Errorf("local store not available. Run 'skool-pp-cli sync' first. Error: %w", err)
			}
			defer db.Close()

			rows, err := db.Query(`
				SELECT id, name, level, points, post_count
				FROM members 
				ORDER BY points DESC
				LIMIT 25
			`)
			if err != nil {
				return err
			}
			defer rows.Close()

			var results []map[string]any
			for rows.Next() {
				var id, name string
				var level, points, postCount int
				if err := rows.Scan(&id, &name, &level, &points, &postCount); err == nil {
					results = append(results, map[string]any{
						"id":         id,
						"name":       name,
						"level":      level,
						"points":     points,
						"post_count": postCount,
					})
				}
			}

			if flags.asJSON {
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(results)
			}

			fmt.Fprintln(cmd.OutOrStdout(), "MEMBER ID\tNAME\tLEVEL\tPOINTS\tPOSTS")
			for _, r := range results {
				fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%d\t%d\t%d\n", r["id"], r["name"], r["level"], r["points"], r["post_count"])
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&period, "period", "7d", "Time window (not currently implemented for dummy data)")
	return cmd
}
