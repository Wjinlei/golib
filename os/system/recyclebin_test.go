package system

import "testing"

func TestRecycleBinStruct_Remove(t *testing.T) {
	bin, err := NewRecycleBin(".")
	if err != nil {
		t.Fatal(err)
	}
	if err := bin.Remove("test_file1.txt"); err != nil {
		t.Fatal(err)
	}
}
