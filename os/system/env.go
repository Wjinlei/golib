package system

import (
	"os"
	"path/filepath"
)

// HostProc 从环境变量获取proc路径
func HostProc(combineWith ...string) string {
	return GetEnv("HOST_PROC", "/proc", combineWith...)
}

// HostSys 从环境变量获取sys路径
func HostSys(combineWith ...string) string {
	return GetEnv("HOST_SYS", "/sys", combineWith...)
}

// HostEtc 从环境变量获取etc路径
func HostEtc(combineWith ...string) string {
	return GetEnv("HOST_ETC", "/etc", combineWith...)
}

// HostVar 从环境变量获取var路径
func HostVar(combineWith ...string) string {
	return GetEnv("HOST_VAR", "/var", combineWith...)
}

// HostRun 从环境变量获取run路径
func HostRun(combineWith ...string) string {
	return GetEnv("HOST_RUN", "/run", combineWith...)
}

// HostDev 从环境变量获取dev路径
func HostDev(combineWith ...string) string {
	return GetEnv("HOST_DEV", "/dev", combineWith...)
}

// HostTmp 从环境变量获取tmp路径
func HostTmp(combineWith ...string) string {
	return GetEnv("HOST_TMP", "/tmp", combineWith...)
}

// GetEnv retrieves the environment variable key. If it does not exist it returns the default.
func GetEnv(key string, defaultValue string, combineWith ...string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	switch len(combineWith) {
	case 0:
		return value
	case 1:
		return filepath.Join(value, combineWith[0])
	default:
		all := make([]string, len(combineWith)+1)
		all[0] = value
		copy(all[1:], combineWith)
		return filepath.Join(all...)
	}
}
