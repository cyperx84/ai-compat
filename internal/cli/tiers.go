package cli

import (
	"fmt"

	"github.com/cyperx/ai-compat/internal/data"
	"github.com/spf13/cobra"
)

type tierPayload struct {
	Kind  string `json:"kind"`
	Items any    `json:"items"`
}

func NewTiersCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tiers",
		Short: "View tier rankings for models or harnesses",
	}

	cmd.AddCommand(newModelTiersCommand())
	cmd.AddCommand(newHarnessTiersCommand())

	return cmd
}

func newModelTiersCommand() *cobra.Command {
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "models",
		Short: "Rank models by aggregate combo score",
		RunE: func(cmd *cobra.Command, args []string) error {
			compat, err := data.LoadData()
			if err != nil {
				return err
			}

			rankings := compat.ModelRankings()
			if jsonOutput {
				return writeJSON(cmd.OutOrStdout(), tierPayload{
					Kind:  "models",
					Items: rankings,
				})
			}

			writeLine(cmd.OutOrStdout(), "Model tiers")
			for _, ranking := range rankings {
				name := ""
				if ranking.Model != nil {
					name = ranking.Model.Name
				}
				writeLine(cmd.OutOrStdout(), "%s  %s (%.1f across %d combos)", ranking.Tier, name, ranking.AggregateScore, ranking.ComboCount)
				if ranking.BestCombo != nil {
					writeLine(cmd.OutOrStdout(), "   Best combo: %s", ranking.BestCombo.Name)
				}
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output as JSON")
	return cmd
}

func newHarnessTiersCommand() *cobra.Command {
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "harnesses",
		Short: "Rank harnesses by aggregate combo score",
		RunE: func(cmd *cobra.Command, args []string) error {
			compat, err := data.LoadData()
			if err != nil {
				return err
			}

			rankings := compat.HarnessRankings()
			if jsonOutput {
				return writeJSON(cmd.OutOrStdout(), tierPayload{
					Kind:  "harnesses",
					Items: rankings,
				})
			}

			writeLine(cmd.OutOrStdout(), "Harness tiers")
			for _, ranking := range rankings {
				name := ""
				if ranking.Harness != nil {
					name = ranking.Harness.Name
				}
				writeLine(cmd.OutOrStdout(), "%s  %s (%.1f across %d combos)", ranking.Tier, name, ranking.AggregateScore, ranking.ComboCount)
				if ranking.BestCombo != nil {
					writeLine(cmd.OutOrStdout(), "   Best combo: %s", ranking.BestCombo.Name)
				}
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output as JSON")
	return cmd
}

func tierLabel(tier data.Tier) string {
	switch tier {
	case data.TierS:
		return "Exceptional"
	case data.TierA:
		return "Great"
	case data.TierB:
		return "Solid"
	case data.TierC:
		return "Usable"
	default:
		return fmt.Sprintf("%s", tier)
	}
}
