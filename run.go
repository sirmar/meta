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
	meta *Meta
	cmd  string
}

func (r *Runner) Run(args []string) {
	r.meta.MoveToRoot()
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

	r.meta.MoveToCwd()
}

func NewDockerRunner(meta *Meta) IRunner {
	return &Runner{meta, "docker"}
}

func NewRunner(meta *Meta, cmd string) IRunner {
	return &Runner{meta, cmd}
}
