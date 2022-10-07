package cmd

import (
	"os"

	"github.com/nao1215/deapk/apk"
	"github.com/nao1215/deapk/internal/print"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info APK_FILES",
	Short: "Print info meta-data from android package (.apk)",
	Long:  `Print info meta-data from android package (.apk)`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := all(cmd, args); err != nil {
			print.Err(err)
			osExit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func all(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return ErrNotSpecifyAPK
	}

	apk := apk.NewAPK(args[0])
	if err := apk.Parse(); err != nil {
		return err
	}
	apk.Print(os.Stdout)
	return nil
}
