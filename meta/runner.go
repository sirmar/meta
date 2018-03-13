package meta

import (
	"os"
	"os/exec"
	"strings"
)

//go:generate mockery -name=IRunner
type IRunner interface {
	Run(cmd string, args []string)
	Output(cmd string, args []string) string
}

type Runner struct {
	log ILog
}

func NewRunner(log ILog) IRunner {
	return &Runner{log}
}

func (self *Runner) Run(command string, args []string) {
	self.log.Println("Running:", command, strings.Join(args, " "))
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		self.log.Fatal(err)
	}
}

func (self *Runner) Output(command string, args []string) string {
	self.log.Println("Running:", command, strings.Join(args, " "))
	output, err := exec.Command(command, args...).Output()
	if err != nil {
		self.log.Fatal(err)
	}
	return strings.TrimSpace(string(output))
}
