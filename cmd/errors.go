package cmd

import (
	"errors"
)

var (
	// ErrNotSpecifyAPK : user does not specify *.apk
	ErrNotSpecifyAPK = errors.New("need to specify one or more apk packages")
)
