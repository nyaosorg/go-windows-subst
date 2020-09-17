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
	return windows.UTF16ToString(targetPath16[:n]), nil
}

func Query(deviceName string) (string, error) {
	return queryDosDevice(deviceName)
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

func List() ([]string, error) {
	return queryDosDevices()
}

const (
	_DDD_EXACT_MATCH_ON_REMOVE = 0x4
	_DDD_NO_BROADCAST_SYSTEM   = 0x8
	_DDD_RAW_TARGET_PATH       = 0x1
	_DDD_REMOVE_DEFINITION     = 0x2
)

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

func Define(deviceName, targetPath string) error {
	return defineDosDevice(0, deviceName, targetPath)
}

func Remove(deviceName string) error {
	return defineDosDevice(_DDD_REMOVE_DEFINITION, deviceName, "")
}
