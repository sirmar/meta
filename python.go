package main

import (
	"fmt"
)

type Python struct {
	Language
}

func NewPython(root *Root, config *Config) *Python {
	srcVolume := fmt.Sprintf("%s:/usr/src/app", root.Root)
	runner := NewDockerRunner(root)
	return &Python{Language{config.Name, config.Name, srcVolume, runner}}
}

func (p *Python) Test() {
	p.dockerRun("nosetests", "test")
}

func (p *Python) Lint() {
	p.dockerRun("flake8", "setup.py", "test", p.name)
}

func (p *Python) Coverage() {
	p.dockerRun("nosetests", "--with-coverage", "test")
}

func (p *Python) CI() {
	p.Install()
	p.Lint()
	p.Coverage()
}

func (p *Python) Run(args []string) {
	p.dockerRun(args...)
}
