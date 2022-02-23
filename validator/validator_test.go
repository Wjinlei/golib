package validator

import "testing"

func TestExists(t *testing.T) {
	exists := Exists("validator.go")
	if exists != true {
		t.Fatal()
	}
}
