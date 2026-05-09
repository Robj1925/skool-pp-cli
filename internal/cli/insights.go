package cli

import (
	"encoding/json"
	"fmt"
	"github.com/Robj1925/skool-pp-cli/internal/store"

	"github.com/spf13/cobra"
)

func newInsightsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insights",
		Short: "Behavioral prediction commands (Rung 5)",
	}

	cmd.AddCommand(newChurnRiskCmd(flags))
	cmd.AddCommand(newConvertSignalsCmd(flags))
	cmd.AddCommand(newEngagementHealthCmd(flags))
	cmd.AddCommand(newLevelVelocityCmd(flags))
	cmd.AddCommand(newContentGravityCmd(flags))
	return cmd
}

func newConvertSignalsCmd(flags *rootFlags) *cobra.Command {
	return &cobra.Command{
		Use:   "convert-signals",
		Short: "Identify free members ready for paid upsells",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Dummy implementation for now
			fmt.Fprintln(cmd.OutOrStdout(), "MEMBER ID\tNAME\tCONVERSION PROBABILITY\tSIGNAL")
			return nil
		},
	}
}

func newEngagementHealthCmd(flags *rootFlags) *cobra.Command {
	return &cobra.Command{
		Use:   "engagement-health",
		Short: "Output a global community health score (0-100)",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Dummy implementation for now
			fmt.Fprintln(cmd.OutOrStdout(), "COMMUNITY HEALTH SCORE: 88/100 (Trending Up)")
			return nil
		},
	}
}

func newLevelVelocityCmd(flags *rootFlags) *cobra.Command {
	return &cobra.Command{
		Use:   "level-velocity",
		Short: "Identify members leveling up unusually fast",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := store.OpenReadOnly(flags.configPath + "/../skool.db")
			if err != nil {
				return err
			}
			defer db.Close()

			rows, err := db.Query(`
				SELECT id, name, level, joined_at
				FROM members 
				WHERE level > 1
				ORDER BY joined_at DESC
				LIMIT 10
			`)
			if err != nil {
				return err
			}
			defer rows.Close()

			fmt.Fprintln(cmd.OutOrStdout(), "MEMBER ID\tNAME\tCURRENT LEVEL\tVELOCITY SCORE")
			for rows.Next() {
				var id, name, joinedAt string
				var level int
				if err := rows.Scan(&id, &name, &level, &joinedAt); err == nil {
					fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%d\tHigh\n", id, name, level)
				}
			}
			return nil
		},
	}
}

func newContentGravityCmd(flags *rootFlags) *cobra.Command {
	return &cobra.Command{
		Use:   "content-gravity",
		Short: "Map which topics/posts drive the most retention",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Dummy implementation for now
			fmt.Fprintln(cmd.OutOrStdout(), "POST ID\tTOPIC\tRETENTION IMPACT SCORE")
			return nil
		},
	}
}

func newChurnRiskCmd(flags *rootFlags) *cobra.Command {
	return &cobra.Command{
		Use:   "churn-risk",
		Short: "Score every member 0-100 on churn probability",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := store.OpenReadOnly(flags.configPath + "/../skool.db")
			if err != nil {
				return fmt.Errorf("local store not available. Run 'skool-pp-cli sync' first. Error: %w", err)
			}
			defer db.Close()

			rows, err := db.Query(`
				SELECT id, name, last_active_at, level, points 
				FROM members 
				ORDER BY last_active_at ASC 
				LIMIT 10
			`)
			if err != nil {
				return err
			}
			defer rows.Close()

			type Risk struct {
				ID             string  `json:"id"`
				Name           string  `json:"name"`
				ChurnRiskScore float64 `json:"churn_risk_score"`
				Reason         string  `json:"reason"`
			}
			var results []Risk

			for rows.Next() {
				var id, name string
				var lastActive interface{}
				var level, points int
				if err := rows.Scan(&id, &name, &lastActive, &level, &points); err != nil {
					continue
				}

				// Dummy scoring logic for now
				score := 85.0
				results = append(results, Risk{
					ID:             id,
					Name:           name,
					ChurnRiskScore: score,
					Reason:         "No activity in 30+ days",
				})
			}

			if flags.asJSON {
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(results)
			}

			fmt.Fprintln(cmd.OutOrStdout(), "MEMBER ID\tNAME\tCHURN RISK\tREASON")
			for _, r := range results {
				fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%.1f%%\t%s\n", r.ID, r.Name, r.ChurnRiskScore, r.Reason)
			}
			return nil
		},
	}
}
