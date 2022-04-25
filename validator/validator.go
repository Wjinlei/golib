package validator

import "os"

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// IsDir 判断所给路径是否是目录
func IsDir(path string) bool {
	stat, err := os.Lstat(path)
	return stat.IsDir()
}
