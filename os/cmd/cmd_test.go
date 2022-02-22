package cmd

import (
	"io/ioutil"
	"testing"
)

func TestRun(t *testing.T) {
	out, err := New().Run("cmd.exe", []string{"/c", "dir"})
	if err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestShell(t *testing.T) {
	out, err := New().Shell("ls")
	if err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestCmd(t *testing.T) {
	out, err := New().Cmd("dir")
	if err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
