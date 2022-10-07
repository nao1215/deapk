// Package apk manage APK file information.
package apk

import (
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
