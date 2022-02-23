package md5

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

func md5sum(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func String(str string) string {
	return md5sum([]byte(str))
}

func File(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return md5sum(data), nil
}
