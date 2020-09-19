package subst

import (
	"fmt"
	"golang.org/x/sys/windows"
)

func utf16toStringArray(s []uint16) []string {
	top := 0
	p := 0
	array := []string{}
	for p < len(s) {
		if s[p] == 0 {
			if p > top {
				tmp := windows.UTF16ToString(s[top:p])
				array = append(array, tmp)
			}
			top = p + 1
		}
		p++
	}
	if p > top {
		tmp := windows.UTF16ToString(s[top:p])
		array = append(array, tmp)
	}
	return array
}

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
	result := windows.UTF16ToString(targetPath16[:n])
	if len(result) >= 4 && result[:4] == `\??\` {
		result = result[4:]
	}
	return result, nil
}

func queryDosDevices() ([]string, error) {
	var targetPath16 [65536]uint16

	n, err := windows.QueryDosDevice(nil,
		&targetPath16[0],
		uint32(len(targetPath16)-1))

	if err != nil {
		return nil, err
	}

	return utf16toStringArray(targetPath16[:n]), nil
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
