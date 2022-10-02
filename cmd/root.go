package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var osExit = os.Exit

var rootCmd = &cobra.Command{
	Use: "apkparser",
}

// Execute start command.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceErrors = true

	if err := rootCmd.Execute(); err != nil {
		osExit(1)
	}
}
