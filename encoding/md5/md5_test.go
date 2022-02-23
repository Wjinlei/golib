package md5

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(Md5([]byte("Hello")))
}

func TestMd5String(t *testing.T) {
	fmt.Println(Md5String("Hello"))
}

func TestMd5File(t *testing.T) {
	md5, err := Md5Path("./md5.go")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(md5)
}
