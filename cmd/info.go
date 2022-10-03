package cmd

import (
	"deapk/internal/print"

	"github.com/shogo82148/androidbinary/apk"
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

	var e error
	for _, v := range args {
		if err := getAllMetadata(v); err != nil {
			// TODO: Use errors.Join() in the future
			print.Warn(err)
			e = ErrNotGetAllMeta
		}
	}
	return e
}

func getAllMetadata(apkPath string) error {
	pkg, err := apk.OpenFile(apkPath)
	if err != nil {
		return err
	}
	defer pkg.Close()

	print.Info(pkg.PackageName())

	return nil
}
