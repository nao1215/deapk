package cmd

import (
	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all APK_FILES",
	Short: "Print all meta-data from android package (.apk)",
	Long:  `Print all meta-data from android package (.apk)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return all(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}

func all(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return ErrNotSpecifyAPK
	}
	return nil
}
