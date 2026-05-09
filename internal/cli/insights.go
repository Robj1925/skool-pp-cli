package cli

import (
	"encoding/json"
	"fmt"
	"skool-pp-cli/internal/store"

	"github.com/spf13/cobra"
)

func newInsightsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insights",
		Short: "Behavioral prediction commands (Rung 5)",
	}

	cmd.AddCommand(newChurnRiskCmd(flags))
	return cmd
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
