package subst

import "strings"

const (
	_DDD_EXACT_MATCH_ON_REMOVE = 0x4
	_DDD_NO_BROADCAST_SYSTEM   = 0x8
	_DDD_RAW_TARGET_PATH       = 0x1
	_DDD_REMOVE_DEFINITION     = 0x2
)

const _PREFIX = `\??\`

func QueryRaw(deviceName string) (string, error) {
	return queryDosDevice(deviceName)
}

func Query(deviceName string) (string, error) {
	result, err := queryDosDevice(deviceName)
	return strings.TrimPrefix(result, _PREFIX), err
}

func Define(deviceName, targetPath string) error {
	return defineDosDevice(0, deviceName, targetPath)
}

func Remove(deviceName string) error {
	return defineDosDevice(_DDD_REMOVE_DEFINITION, deviceName, "")
}
