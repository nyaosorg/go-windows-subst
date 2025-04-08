//go:build !windows
// +build !windows

package subst

import (
	"errors"
)

func queryDosDevice(string) (string, error) {
	return "", errors.New("not supported in UNIX")
}

func defineDosDevice(_ uint32, _ string, _ string) error {
	return errors.New("not supported in UNIX")
}
