package subst

import "strings"

const (
	_DDD_EXACT_MATCH_ON_REMOVE = 0x4
	_DDD_NO_BROADCAST_SYSTEM   = 0x8
	_DDD_RAW_TARGET_PATH       = 0x1
	_DDD_REMOVE_DEFINITION     = 0x2
)

const _PREFIX = `\??\`

// QueryRaw returns the path assigned to a drive letter via subst.exe or similar tools.
// Unlike Query, it does not remove the `\??\` prefix returned by the Windows API.
func QueryRaw(deviceName string) (string, error) {
	return queryDosDevice(deviceName)
}

// Query returns the path assigned to a drive letter via subst.exe or similar tools.
// It uses the Windows API QueryDosDevice to obtain the path, and removes the `\??\` prefix if present.
func Query(deviceName string) (string, error) {
	result, err := queryDosDevice(deviceName)
	return strings.TrimPrefix(result, _PREFIX), err
}

// Define assigns an existing path to a drive letter, similar to the subst command.
func Define(deviceName, targetPath string) error {
	return defineDosDevice(0, deviceName, targetPath)
}

// Remove unmaps a drive letter previously defined by subst.
func Remove(deviceName string) error {
	return defineDosDevice(_DDD_REMOVE_DEFINITION, deviceName, "")
}
