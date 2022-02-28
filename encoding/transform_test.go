package encoding

import (
	"fmt"
	"testing"
)

func TestGbkToUtf8(t *testing.T) {
	utf8, err := GbkToUtf8([]byte("你好世界"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(utf8))
}

func TestUtf8ToGbk(t *testing.T) {
	gbk, err := Utf8ToGbk([]byte("你好世界"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(gbk))
}
