package validator

import "os"

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func Has[T comparable](a []T, b T) bool {
	for _, element := range a {
		if element == b {
			return true
		}
	}
	return false
}
