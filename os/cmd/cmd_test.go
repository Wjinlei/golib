package cmd

import (
	"io/ioutil"
	"testing"
)

func TestCmdRunner_Run(t *testing.T) {
	out, err := New().Run("cmd.exe", []string{"/c", "dir"})
	if err != nil {
		t.Fatal(err)
	}
	_, err = ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCmdRunner_Shell(t *testing.T) {
	out, err := New().Shell("ls")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCmdRunner_Cmd(t *testing.T) {
	out, err := New().Cmd("dir")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}
}
