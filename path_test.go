package golib

import (
	"fmt"
	"testing"
)

func TestGetAbsPath(t *testing.T) {
	abs, err := GetAbsPath("os/cmd")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(abs)
}
