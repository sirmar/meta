package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type IRunner interface {
	Run(args []string)
}

type Runner struct {
	root *Root
	cmd  string
}

func (r *Runner) Run(args []string) {
	r.root.MoveToRoot()
	fmt.Println("Running:", r.cmd, strings.Join(args, " "))

	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command(r.cmd, args...)

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}

	r.root.MoveToCwd()
}

func NewDockerRunner(root *Root) IRunner {
	return &Runner{root, "docker"}
}

func NewRunner(root *Root, cmd string) IRunner {
	return &Runner{root, cmd}
}
