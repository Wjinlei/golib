package golib

import (
	"fmt"
	"path/filepath"
)

// GetAbs 返回传入路径的绝对路径，如果获取失败则返回原路径
func GetAbs(oldPath string) string {
	if filepath.IsAbs(oldPath) {
		return oldPath
	}
	newPath, err := filepath.Abs(oldPath)
	if err != nil {
		return oldPath
	}
	return filepath.FromSlash(fmt.Sprintf("%s", newPath))
}
