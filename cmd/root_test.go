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
		orgHome := os.Getenv("HOME")
		if err := os.Setenv("HOME", t.TempDir()); err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := os.Setenv("HOME", orgHome); err != nil {
				t.Fatal(err)
			}
		}()

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

func TestExecute_NotExistSubCommand(t *testing.T) {
	t.Run("non-existent subcommand is specified", func(t *testing.T) {
		exitCode := 0
		osExit = func(code int) {
			exitCode = code
		}
		defer func() { osExit = os.Exit }()

		os.Args = []string{"deapk", "non-exist-sub-command"}
		Execute()

		want := 1
		if want != exitCode {
			t.Errorf("mismatch want=%d, got=%d", want, exitCode)
		}
	})

	t.Run("output metadata at stdout in default formtat", func(t *testing.T) {
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
	})

	t.Run("output metadata at stdout in json formtat", func(t *testing.T) {
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
	})
}

func TestExecute_Version(t *testing.T) {
	t.Run("success to get version", func(t *testing.T) {
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
		Version = "v1.0.0"
		os.Args = []string{"deapk", "version"}
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
			"deapk version v1.0.0",
			"",
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("failed to get version", func(t *testing.T) {
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
		Version = ""
		os.Args = []string{"deapk", "version"}
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
			"deapk version ",
			"",
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}
