package golib

import (
	"testing"
)

func TestMkdir(t *testing.T) {
	if err := MakeDir("test_dir1/sub1/sub2/sub3"); err != nil {
		t.Fatal(err)
	}
}

func TestMakeDirParent(t *testing.T) {
	if err := MakeDirParent("test_dir2/sub1/sub2/sub3"); err != nil {
		t.Fatal(err)
	}
}

func TestFileWrite(t *testing.T) {
	if err := FileWrite("test_file1.txt", "Hello World\r\n", FileCreate); err != nil {
		t.Fatal(err)
	}
}

func TestFileAppend(t *testing.T) {
	if err := FileWrite("test_file1.txt", "Hello World\r\n", FileAppend); err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	if err := Delete("test_dir1"); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteAll(t *testing.T) {
	if err := DeleteAll("test_dir1"); err != nil {
		t.Fatal(err)
	}
}

func TestMove(t *testing.T) {
	if err := Move("test_dir1", "test_dir2"); err != nil {
		t.Fatal(err)
	}
}

func TestReadFile(t *testing.T) {
	_, err := ReadFile("test_file1.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadLines(t *testing.T) {
	_, err := ReadLines("test_file1.txt", "\r\n")
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadLinesOffsetN(t *testing.T) {
	_, err := ReadLinesOffsetN("test_file1.txt", 0, 2, "\r\n")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCopy(t *testing.T) {
	if err := Copy("test_file1.txt", "test_file2.txt"); err != nil {
		t.Fatal(err)
	}
	if err := Copy("test_dir1", "test_dir2"); err != nil {
		t.Fatal(err)
	}
}

func TestFileDownload(t *testing.T) {
	err := FileDownload("http://softdown.huweishen.com/4/VCRedistInstaller2019.zip", "VCRedistInstaller2019.zip")
	if err != nil {
		t.Fatal(err)
	}
}
