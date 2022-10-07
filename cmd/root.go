package cmd

import (
	"os"

	"github.com/nao1215/deapk/internal/print"

	"github.com/spf13/cobra"
)

var osExit = os.Exit

var rootCmd = &cobra.Command{
	Use: "deapk",
}

// Execute start command.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceErrors = true

	if err := rootCmd.Execute(); err != nil {
		print.Err(err)
		osExit(1)
	}
}
