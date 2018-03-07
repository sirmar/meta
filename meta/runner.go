package meta

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
)

//go:generate mockery -name=IRunner
type IRunner interface {
	Run(cmd string, args []string)
}

type Runner struct {
	log ILog
}

func NewRunner(log ILog) IRunner {
	return &Runner{log}
}

func (self *Runner) Run(command string, args []string) {
	self.log.Println("Running:", command, strings.Join(args, " "))

	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command(command, args...)

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		self.log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		self.log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		self.log.Fatal("failed to capture stdout or stderr\n")
	}
}
