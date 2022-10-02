package cmd

import (
	"errors"
	"testing"
)

func Test_wrap(t *testing.T) {
	t.Run("success to wrap error interface", func(t *testing.T) {
		want := "new message: original message"
		got := wrap("new message", errors.New("original message"))

		if want != got.Error() {
			t.Errorf("mismatch: want=%s, got=%v", want, got)
		}
	})
}
