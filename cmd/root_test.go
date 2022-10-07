// Package cmd define subcommands provided by the apk-parser command
package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nao1215/morrigan/file"
)

func TestExecute_Completion(t *testing.T) {
	t.Run("generate completion file", func(t *testing.T) {
		os.Args = []string{"deapk", "completion"}
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

func TestExecute_OutputMetadata(t *testing.T) {
	orgStdout := os.Stdout
	orgStderr := os.Stderr
	pr, pw, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = pw
	os.Stderr = pw

	osExit = func(code int) {}
	defer func() {
		osExit = os.Exit
	}()

	// Test start
	os.Args = []string{"deapk", "info", "../testdata/app-debug.apk"}
	Execute()

	pw.Close()
	os.Stdout = orgStdout
	os.Stderr = orgStderr

	buf := bytes.Buffer{}
	_, err = io.Copy(&buf, pr)
	if err != nil {
		t.Error(err)
	}
	defer pr.Close()
	got := strings.Split(buf.String(), "\n")

	want := []string{
		"pacakage name      : jp.debimate.deapk_test",
		"application name   : deapk-test",
		"application version: 1.0",
		"sdk target version : 31",
		"sdk max version    : -1 (deprecated attribute)",
		"sdk min version    : 31",
		"main activity      : jp.debimate.deapk_test.MainActivity",
		"",
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}
}

func TestExecute_OutputMetadataInJson(t *testing.T) {
	orgStdout := os.Stdout
	orgStderr := os.Stderr
	pr, pw, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = pw
	os.Stderr = pw

	osExit = func(code int) {}
	defer func() {
		osExit = os.Exit
	}()

	// Test start
	os.Args = []string{"deapk", "info", "--json", "../testdata/app-debug.apk"}
	Execute()

	pw.Close()
	os.Stdout = orgStdout
	os.Stderr = orgStderr

	buf := bytes.Buffer{}
	_, err = io.Copy(&buf, pr)
	if err != nil {
		t.Error(err)
	}
	defer pr.Close()
	got := strings.Split(buf.String(), "\n")

	want := []string{
		"{",
		`	"Basic": {`,
		`		"package_name": "jp.debimate.deapk_test",`,
		`		"application_name": "deapk-test",`,
		`		"version": "1.0",`,
		`		"main_activity": "jp.debimate.deapk_test.MainActivity",`,
		`		"sdk": {`,
		`			"minimum": 31,`,
		`			"target": 31,`,
		`			"maximum": -1`,
		`		}`,
		`	}`,
		"}",
		"",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}
}
