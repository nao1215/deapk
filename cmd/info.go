package cmd

import (
	"fmt"
	"os"

	"github.com/nao1215/deapk/apk"
	"github.com/nao1215/deapk/internal/print"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info APK_FILES",
	Short: "Print meta-data for android package (.apk)",
	Long:  `Print meta-data for android package (.apk)`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := all(cmd, args); err != nil {
			print.Err(err)
			osExit(1)
		}
	},
}

func init() {
	infoCmd.Flags().StringP("output", "o", "", "output apk information to the file")
	rootCmd.AddCommand(infoCmd)
}

func all(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return ErrNotSpecifyAPK
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return (fmt.Errorf("%s: %w", "can not parse command line argument (--output)", err))
	}

	apk := apk.NewAPK(args[0])
	if err := apk.Parse(); err != nil {
		return err
	}

	writer := os.Stdout
	if output != "" {
		f, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return nil
		}
		defer f.Close()
		writer = f
	}
	apk.Print(writer)

	return nil
}
