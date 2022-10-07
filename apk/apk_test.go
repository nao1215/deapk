// Package apk manage APK file information.
package apk

import (
	"errors"
	"fmt"
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
}
