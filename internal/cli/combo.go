package cli

import (
	"fmt"

	"github.com/cyperx/ai-compat/internal/data"
	"github.com/spf13/cobra"
)

type comboPayload struct {
	Combo   *data.Combo   `json:"combo"`
	Model   *data.Model   `json:"model"`
	Harness *data.Harness `json:"harness"`
	Usecase *data.Usecase `json:"usecase,omitempty"`
}

func NewComboCommand() *cobra.Command {
	var modelSlug string
	var harnessSlug string
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "combo --model <slug> --harness <slug>",
		Short: "Inspect one model and harness combination",
		RunE: func(cmd *cobra.Command, args []string) error {
			if modelSlug == "" || harnessSlug == "" {
				return fmt.Errorf("both --model and --harness are required")
			}

			compat, err := data.LoadData()
			if err != nil {
				return err
			}

			combo := compat.FindComboByParts(modelSlug, harnessSlug)
			if combo == nil {
				return fmt.Errorf("no combo found for %s + %s", modelSlug, harnessSlug)
			}

			payload := comboPayload{
				Combo:   combo,
				Model:   compat.FindModel(modelSlug),
				Harness: compat.FindHarness(harnessSlug),
			}
			if combo.Usecase != "" {
				payload.Usecase = compat.FindUsecase(combo.Usecase)
			}

			if jsonOutput {
				return writeJSON(cmd.OutOrStdout(), payload)
			}

			writeLine(cmd.OutOrStdout(), "%s", combo.Name)
			writeLine(cmd.OutOrStdout(), "Score: %.1f/10", combo.Score)
			writeLine(cmd.OutOrStdout(), "Status: %s", combo.Status)
			if payload.Model != nil {
				writeLine(cmd.OutOrStdout(), "Model: %s", payload.Model.Name)
			}
			if payload.Harness != nil {
				writeLine(cmd.OutOrStdout(), "Harness: %s", payload.Harness.Name)
			}
			if payload.Usecase != nil {
				writeLine(cmd.OutOrStdout(), "Best for: %s", payload.Usecase.Name)
			}
			if combo.Notes != "" {
				writeLine(cmd.OutOrStdout(), "Notes: %s", combo.Notes)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&modelSlug, "model", "", "Model slug")
	cmd.Flags().StringVar(&harnessSlug, "harness", "", "Harness slug")
	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output as JSON")
	return cmd
}
