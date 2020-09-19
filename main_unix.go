// +build !windows

package subst

import (
	"errors"
)

func queryDosDevices() ([]string, error) {
	return nil, errors.New("not supported in UNIX")
}

func queryDosDevice(string) (string, error) {
	return "", errors.New("not supported in UNIX")
}

func defineDosDevice(_ uint32, _ string, _ string) error {
	return errors.New("not supported in UNIX")
}
