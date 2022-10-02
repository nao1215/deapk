// Package cmd define subcommands provided by the apk-parser command
package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/nao1215/morrigan/file"
)

func TestExecute_Completion(t *testing.T) {
	t.Run("generate completion file", func(t *testing.T) {
		os.Args = []string{"gup", "completion"}
		Execute()

		bash := filepath.Join(os.Getenv("HOME"), ".bash_completion")
		if runtime.GOOS == "windows" {
			if file.IsFile(bash) {
				t.Errorf("generate %s, however shell completion file is not generated on Windows", bash)
			}
		} else {
			if !file.IsFile(bash) {
				t.Errorf("failed to generate %s", bash)
			}
		}

		fish := filepath.Join(os.Getenv("HOME"), ".config", "fish", "completions", Name+".fish")
		if runtime.GOOS == "windows" {
			if file.IsFile(fish) {
				t.Errorf("generate %s, however shell completion file is not generated on Windows", fish)
			}
		} else {
			if !file.IsFile(fish) {
				t.Errorf("failed to generate %s", fish)
			}
		}

		zsh := filepath.Join(os.Getenv("HOME"), ".zsh", "completion", "_"+Name)
		if runtime.GOOS == "windows" {
			if file.IsFile(zsh) {
				t.Errorf("generate %s, however shell completion file is not generated on Windows", zsh)
			}
		} else {
			if !file.IsFile(zsh) {
				t.Errorf("failed to generate  %s", zsh)
			}
		}
	})
}

func TestExecute(t *testing.T) {
	t.Run("not specify apk", func(t *testing.T) {
		os.Args = []string{"apkparser"}
		want := ErrNotSpecifyAPK
		got := rootCmd.Execute()

		if errors.Is(got, want) {
			t.Errorf("mismatch: want=%v, got=%v", want, got)
		}
	})
}
