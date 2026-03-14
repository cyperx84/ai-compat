package cli

import (
	"fmt"
	"strings"

	"github.com/cyperx/ai-compat/internal/data"
	"github.com/spf13/cobra"
)

type comparePayload struct {
	Kind  string `json:"kind"`
	Left  any    `json:"left"`
	Right any    `json:"right"`
}

func NewCompareCommand() *cobra.Command {
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "compare <slugA> <slugB>",
		Short: "Compare two models or two harnesses",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			compat, err := data.LoadData()
			if err != nil {
				return err
			}

			leftModel := compat.FindModel(args[0])
			rightModel := compat.FindModel(args[1])
			if leftModel != nil && rightModel != nil {
				return compareModels(cmd, leftModel, rightModel, jsonOutput)
			}

			leftHarness := compat.FindHarness(args[0])
			rightHarness := compat.FindHarness(args[1])
			if leftHarness != nil && rightHarness != nil {
				return compareHarnesses(cmd, leftHarness, rightHarness, jsonOutput)
			}

			return fmt.Errorf("compare expects two model slugs or two harness slugs")
		},
	}

	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output as JSON")
	return cmd
}

func compareModels(cmd *cobra.Command, left, right *data.Model, jsonOutput bool) error {
	if jsonOutput {
		return writeJSON(cmd.OutOrStdout(), comparePayload{
			Kind:  "model",
			Left:  left,
			Right: right,
		})
	}

	writeLine(cmd.OutOrStdout(), "Model comparison")
	writeLine(cmd.OutOrStdout(), "%s vs %s", left.Name, right.Name)
	writeLine(cmd.OutOrStdout(), "Providers: %s | %s", left.Provider, right.Provider)
	writeLine(cmd.OutOrStdout(), "Context: %s | %s", left.ContextWindow, right.ContextWindow)
	writeLine(cmd.OutOrStdout(), "Released: %s | %s", left.Released, right.Released)
	writeLine(cmd.OutOrStdout(), "Capabilities: %s | %s",
		strings.Join(left.Capabilities, ", "),
		strings.Join(right.Capabilities, ", "),
	)
	return nil
}

func compareHarnesses(cmd *cobra.Command, left, right *data.Harness, jsonOutput bool) error {
	if jsonOutput {
		return writeJSON(cmd.OutOrStdout(), comparePayload{
			Kind:  "harness",
			Left:  left,
			Right: right,
		})
	}

	writeLine(cmd.OutOrStdout(), "Harness comparison")
	writeLine(cmd.OutOrStdout(), "%s vs %s", left.Name, right.Name)
	writeLine(cmd.OutOrStdout(), "Providers: %s | %s", left.Provider, right.Provider)
	writeLine(cmd.OutOrStdout(), "Types: %s | %s", left.Type, right.Type)
	writeLine(cmd.OutOrStdout(), "Status: %s | %s", left.Status, right.Status)
	writeLine(cmd.OutOrStdout(), "Features: %s | %s",
		strings.Join(left.Features, ", "),
		strings.Join(right.Features, ", "),
	)
	return nil
}
