package cmd

import (
	"github.com/nao1215/deapk/internal/completion"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Create shell completion files (bash, fish, zsh) for the deapk",
	Long: `Create shell completion files (bash, fish, zsh) for the deapk command
if it is not already on the system`,
	RunE: func(cmd *cobra.Command, args []string) error {
		completion.DeployShellCompletionFileIfNeeded(rootCmd)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
