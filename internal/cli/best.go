package cli

import (
	"fmt"
	"sort"

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

			var matches []data.Combo
			for _, combo := range compat.Combos {
				if usecase == "" || combo.Usecase == usecase {
					matches = append(matches, combo)
				}
			}

			if len(matches) == 0 {
				return fmt.Errorf("no combos found for use case %q", usecase)
			}

			sort.Slice(matches, func(i, j int) bool {
				return matches[i].Score > matches[j].Score
			})

			limit := 5
			if len(matches) < limit {
				limit = len(matches)
			}

			payload := bestPayload{
				For:    usecase,
				Combos: make([]bestCombo, 0, limit),
			}

			for _, combo := range matches[:limit] {
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
