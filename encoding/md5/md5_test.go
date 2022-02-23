package md5

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println(String("Hello"))
}

func TestFile(t *testing.T) {
	md5, err := File("./md5.go")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(md5)
}
