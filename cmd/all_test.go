package cmd

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
)

func Test_all(t *testing.T) {
	t.Run("not specify apk", func(t *testing.T) {
		want := ErrNotSpecifyAPK
		got := all(&cobra.Command{}, []string{})
		if !errors.Is(got, want) {
			t.Errorf("mismatch: want=%v, got=%v", want, got)
		}
	})
}
