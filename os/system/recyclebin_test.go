package system

import "testing"

func TestRecycleBinStruct_Remove(t *testing.T) {
	bin, err := NewRecycleBin(".")
	if err != nil {
		t.Fatal(err)
	}
	remove, err := bin.Remove("test_file1.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(remove)
}

func TestRecycleBinStruct_Restore(t *testing.T) {
	bin, err := NewRecycleBin(".")
	if err != nil {
		t.Fatal(err)
	}
	remove, err := bin.Remove("test_file1.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = bin.Restore(remove)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecycleBinStruct_Delete(t *testing.T) {
	bin, err := NewRecycleBin(".")
	if err != nil {
		t.Fatal(err)
	}
	remove, err := bin.Remove("test_file1.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = bin.Delete(remove)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecycleBinStruct_Empty(t *testing.T) {
	bin, err := NewRecycleBin(".")
	if err != nil {
		t.Fatal(err)
	}
	_, err = bin.Remove("test_file1.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = bin.Empty()
	if err != nil {
		t.Fatal(err)
	}
}
