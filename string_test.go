package golib

import "testing"

func TestCreateRandomString(t *testing.T) {
	strLen := 8
	str := CreateRandomString(strLen)
	if len(str) != strLen {
		t.Fatalf("string len != %d", strLen)
	}
	t.Log(str)
}
