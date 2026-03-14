package cli

import (
	"strings"

	"github.com/cyperx/ai-compat/internal/data"
	"github.com/spf13/cobra"
)

type searchResult struct {
	Type  string `json:"type"`
	Slug  string `json:"slug"`
	Name  string `json:"name"`
	Match string `json:"match"`
}

func NewSearchCommand() *cobra.Command {
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "search <query>",
		Short: "Search for models, harnesses, or combos",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			compat, err := data.LoadData()
			if err != nil {
				return err
			}

			query := strings.ToLower(args[0])
			results := make([]searchResult, 0)

			for _, model := range compat.Models {
				if matchesModel(model, query) {
					results = append(results, searchResult{
						Type:  "model",
						Slug:  model.Slug,
						Name:  model.Name,
						Match: model.Provider,
					})
				}
			}

			for _, harness := range compat.Harnesses {
				if matchesHarness(harness, query) {
					results = append(results, searchResult{
						Type:  "harness",
						Slug:  harness.Slug,
						Name:  harness.Name,
						Match: harness.Type,
					})
				}
			}

			for _, combo := range compat.Combos {
				if matchesCombo(combo, query) {
					results = append(results, searchResult{
						Type:  "combo",
						Slug:  combo.Slug,
						Name:  combo.Name,
						Match: combo.Usecase,
					})
				}
			}

			if jsonOutput {
				return writeJSON(cmd.OutOrStdout(), results)
			}

			if len(results) == 0 {
				writeLine(cmd.OutOrStdout(), "No results found for %q.", args[0])
				return nil
			}

			for _, result := range results {
				if result.Match != "" {
					writeLine(cmd.OutOrStdout(), "%s  %s (%s) [%s]", strings.ToUpper(result.Type), result.Name, result.Slug, result.Match)
					continue
				}
				writeLine(cmd.OutOrStdout(), "%s  %s (%s)", strings.ToUpper(result.Type), result.Name, result.Slug)
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output as JSON")
	return cmd
}

func matchesModel(model data.Model, query string) bool {
	if strings.Contains(strings.ToLower(model.Name), query) ||
		strings.Contains(strings.ToLower(model.Provider), query) ||
		strings.Contains(strings.ToLower(model.Description), query) ||
		strings.Contains(strings.ToLower(model.Slug), query) {
		return true
	}

	for _, capability := range model.Capabilities {
		if strings.Contains(strings.ToLower(capability), query) {
			return true
		}
	}

	return false
}

func matchesHarness(harness data.Harness, query string) bool {
	if strings.Contains(strings.ToLower(harness.Name), query) ||
		strings.Contains(strings.ToLower(harness.Type), query) ||
		strings.Contains(strings.ToLower(harness.Provider), query) ||
		strings.Contains(strings.ToLower(harness.Description), query) ||
		strings.Contains(strings.ToLower(harness.Slug), query) {
		return true
	}

	for _, feature := range harness.Features {
		if strings.Contains(strings.ToLower(feature), query) {
			return true
		}
	}

	return false
}

func matchesCombo(combo data.Combo, query string) bool {
	return strings.Contains(strings.ToLower(combo.Name), query) ||
		strings.Contains(strings.ToLower(combo.Description), query) ||
		strings.Contains(strings.ToLower(combo.Slug), query) ||
		strings.Contains(strings.ToLower(combo.Notes), query) ||
		strings.Contains(strings.ToLower(combo.Usecase), query)
}
