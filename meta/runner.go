package meta

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
)

type IRunner interface {
	Run(args []string)
}

type Runner struct {
	cmd string
	log ILog
}

func (r *Runner) Run(args []string) {
	r.log.Println("Running:", r.cmd, strings.Join(args, " "))

	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command(r.cmd, args...)

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		r.log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		r.log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		r.log.Fatal("failed to capture stdout or stderr\n")
	}
}

func NewDockerRunner(log ILog) IRunner {
	return NewRunner("docker", log)
}

func NewRunner(cmd string, log ILog) IRunner {
	return &Runner{cmd, log}
}
