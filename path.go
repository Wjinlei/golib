package golib

import (
	"fmt"
	"path/filepath"
	"strings"
)

// GetAbsPath 返回传入路径的绝对路径
func GetAbsPath(filePath string) (string, error) {
	if filepath.IsAbs(filePath) {
		return filePath, nil
	}
	if strings.HasSuffix(filePath, ".") || strings.HasSuffix(filePath, "..") {
		filePath = filePath + "/"
	}
	filePathAbs, err := filepath.Abs(filepath.Dir(filePath))
	if err != nil {
		return "", err
	}
	name := filepath.Base(filePath)
	if name == "." || name == ".." {
		return filepath.FromSlash(filePathAbs), nil
	}
	return filepath.FromSlash(fmt.Sprintf("%s/%s", filePathAbs, filepath.Base(filePath))), nil
}
