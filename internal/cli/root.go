package cli

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aicomp",
		Short: "AI model + harness compatibility CLI",
		Long: `aicomp - Find the best model + harness combination for your workflow.
Compare Claude, GPT, Gemini and others with Claude Code, Codex CLI, OpenClaw, and more.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(NewSearchCommand())
	cmd.AddCommand(NewCompareCommand())
	cmd.AddCommand(NewComboCommand())
	cmd.AddCommand(NewBestCommand())
	cmd.AddCommand(NewTiersCommand())

	return cmd
}
