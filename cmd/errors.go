package cmd

import (
	"errors"
	"fmt"
)

var (
	// ErrNotSpecifyAPK : user does not specify *.apk
	ErrNotSpecifyAPK = errors.New("need to specify one or more apk packages")
	// ErrNotGetAllMeta : failed to get all metadata
	ErrNotGetAllMeta = errors.New("failed to get all metadata")
)

// wrap wrap error interface with new message.
func wrap(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}
