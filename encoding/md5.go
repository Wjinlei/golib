package encoding

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

func Md5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func Md5String(str string) string {
	return Md5([]byte(str))
}

func Md5Path(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return Md5(data), nil
}
