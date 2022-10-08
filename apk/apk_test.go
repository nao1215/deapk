// Package apk manage APK file information.
package apk

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewAPK(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *APK
	}{
		{
			name: "Generate new apk file",
			args: args{
				path: "path/to/apk",
			},
			want: &APK{
				Path: "path/to/apk",
				Package: &Package{
					Basic: &Basic{
						SDK: &SDK{},
					},
					Metadata: []*Metadata{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAPK(tt.args.path)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAPK_Parse(t *testing.T) {
	t.Run("Failed to open apk file", func(t *testing.T) {
		a := &APK{
			Path: "not/exist/path",
		}

		want := fmt.Errorf("%w: %s", ErrNotOpenAPK, a.Path)
		got := a.Parse()
		if errors.Is(want, got) {
			t.Errorf("mismatch want=%v, got=%v", want, got)
		}
	})

	t.Run("Success to get apk information", func(t *testing.T) {
		orgStdout := os.Stdout
		orgStderr := os.Stderr
		pr, pw, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}
		os.Stdout = pw
		os.Stderr = pw

		// Test start
		a := NewAPK("../testdata/app-debug.apk")
		if err := a.Parse(); err != nil {
			t.Fatal(err)
		}
		a.Print(os.Stdout)

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
			"[Application]",
			" name           : deapk-test",
			" version        : 1.0",
			" main activity  : jp.debimate.deapk_test.MainActivity",
			" package        : jp.debimate.deapk_test",
			"[SDK]",
			" target version : 31",
			" max version    : -1 (deprecated attribute)",
			" min version    : 31",
			"",
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}
