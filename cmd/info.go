package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/nao1215/deapk/internal/print"

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
		meta, err := parseMetadata(v)
		if err != nil {
			// TODO: Use errors.Join() in the future
			print.Warn(err)
			e = ErrNotGetAllMeta
		}
		meta.print(os.Stdout)
	}
	return e
}

type metadata struct {
	packageName string
}

func parseMetadata(apkPath string) (*metadata, error) {
	pkg, err := apk.OpenFile(apkPath)
	if err != nil {
		return nil, err
	}
	defer pkg.Close()

	meta := &metadata{}
	meta.packageName = pkg.PackageName()

	return meta, nil
}

func (m *metadata) print(w io.Writer) {
	fmt.Fprintf(w, "pacakage name: %s\n", m.packageName)
}
