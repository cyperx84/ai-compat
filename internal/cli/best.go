package cli

import (
	"fmt"

	"github.com/cyperx/ai-compat/internal/data"
	"github.com/spf13/cobra"
)

type bestCombo struct {
	Combo   data.Combo    `json:"combo"`
	Model   *data.Model   `json:"model"`
	Harness *data.Harness `json:"harness"`
	Usecase *data.Usecase `json:"usecase,omitempty"`
}

type bestPayload struct {
	For    string      `json:"for"`
	Combos []bestCombo `json:"combos"`
}

func NewBestCommand() *cobra.Command {
	var usecase string
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "best",
		Short: "Return top combinations overall or for a use case",
		RunE: func(cmd *cobra.Command, args []string) error {
			compat, err := data.LoadData()
			if err != nil {
				return err
			}

			matches := compat.BestCombos(usecase, 5)

			if len(matches) == 0 {
				return fmt.Errorf("no combos found for use case %q", usecase)
			}

			payload := bestPayload{
				For:    usecase,
				Combos: make([]bestCombo, 0, len(matches)),
			}

			for _, combo := range matches {
				item := bestCombo{
					Combo:   combo,
					Model:   compat.FindModel(combo.Model),
					Harness: compat.FindHarness(combo.Harness),
				}
				if combo.Usecase != "" {
					item.Usecase = compat.FindUsecase(combo.Usecase)
				}
				payload.Combos = append(payload.Combos, item)
			}

			if jsonOutput {
				return writeJSON(cmd.OutOrStdout(), payload)
			}

			if usecase == "" {
				writeLine(cmd.OutOrStdout(), "Top combos")
			} else {
				writeLine(cmd.OutOrStdout(), "Top combos for %s", usecase)
			}
			for index, item := range payload.Combos {
				writeLine(cmd.OutOrStdout(), "%d. %s (%.1f)", index+1, item.Combo.Name, item.Combo.Score)
				if item.Usecase != nil {
					writeLine(cmd.OutOrStdout(), "   %s", item.Usecase.Name)
				}
				if len(item.Combo.Pros) > 0 {
					writeLine(cmd.OutOrStdout(), "   Pros: %s", joinList(item.Combo.Pros))
				}
				if item.Combo.Notes != "" {
					writeLine(cmd.OutOrStdout(), "   %s", item.Combo.Notes)
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&usecase, "for", "", "Use case slug")
	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output as JSON")
	return cmd
}
