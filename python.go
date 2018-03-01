package main

import (
	"fmt"
)

type Python struct {
	Language
}

func NewPython(runner IRunner, meta *Meta) *Python {
	name := meta.Meta.Name
	srcVolume := fmt.Sprintf("%s:/usr/src/%s", meta.Root, name)
	return &Python{Language{name, name, srcVolume, runner}}
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
