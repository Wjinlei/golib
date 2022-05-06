package golib

import (
	"fmt"
	"testing"
)

func TestGetAbsPath(t *testing.T) {
	newPath := GetAbs("os/cmd")
	if newPath == "os/cmd" {
		t.Fatal("GetAbs fatal.")
	}
	fmt.Println(newPath)
}
