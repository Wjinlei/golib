package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"time"
)

type cmdRunner struct{}

func New() *cmdRunner {
	return &cmdRunner{}
}

func (c *cmdRunner) Run(cmd string, args []string) (io.Reader, error) {
	command := exec.Command(cmd, args...)
	resChannel := make(chan []byte)
	errChannel := make(chan error)

	go func() {
		out, err := command.CombinedOutput()
		if err != nil {
			errChannel <- err
		}
		resChannel <- out
	}()

	timer := time.After(5 * time.Second)
	select {
	case err := <-errChannel:
		return nil, err
	case res := <-resChannel:
		return bytes.NewReader(res), nil
	case <-timer:
		return nil, fmt.Errorf("time out (cmd:%v args:%v)", cmd, args)
	}
}

func (c *cmdRunner) Shell(cmd string) (io.Reader, error) {
	command := exec.Command("sh", "-c", cmd)
	resChannel := make(chan []byte)
	errChannel := make(chan error)

	go func() {
		out, err := command.CombinedOutput()
		if err != nil {
			errChannel <- err
		}
		resChannel <- out
	}()

	timer := time.After(5 * time.Second)
	select {
	case err := <-errChannel:
		return nil, err
	case res := <-resChannel:
		return bytes.NewReader(res), nil
	case <-timer:
		return nil, fmt.Errorf("time out (cmd:%v)", cmd)
	}
}

func (c *cmdRunner) Cmd(cmd string) (io.Reader, error) {
	command := exec.Command("cmd.exe", "/c", cmd)
	resChannel := make(chan []byte)
	errChannel := make(chan error)

	go func() {
		out, err := command.CombinedOutput()
		if err != nil {
			errChannel <- err
		}
		resChannel <- out
	}()

	timer := time.After(5 * time.Second)
	select {
	case err := <-errChannel:
		return nil, err
	case res := <-resChannel:
		return bytes.NewReader(res), nil
	case <-timer:
		return nil, fmt.Errorf("time out (cmd:%v)", cmd)
	}
}
