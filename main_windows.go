package subst

import (
	"fmt"
	"golang.org/x/sys/windows"
)

func queryDosDevice(deviceName string) (string, error) {
	deviceNamePtr, err := windows.UTF16PtrFromString(deviceName)
	if err != nil {
		return "", fmt.Errorf("%s: %s", deviceName, err)
	}

	var targetPath16 [1024]uint16

	n, err := windows.QueryDosDevice(deviceNamePtr,
		&targetPath16[0],
		uint32(len(targetPath16)-1))

	if err != nil {
		return "", err
	}
	return windows.UTF16ToString(targetPath16[:n]), nil
}

func defineDosDevice(flags uint32, deviceName string, targetPath string) error {
	deviceNamePtr, err := windows.UTF16PtrFromString(deviceName)
	if err != nil {
		return err
	}
	var targetPathPtr *uint16 = nil
	if targetPath != "" {
		targetPathPtr, err = windows.UTF16PtrFromString(targetPath)
		if err != nil {
			return err
		}
	}
	return windows.DefineDosDevice(flags, deviceNamePtr, targetPathPtr)
}
