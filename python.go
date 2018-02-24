package main

import (
	"fmt"
	"log"
)

type Python struct {
	name      string
	image     string
	srcVolume string
	runner    *Runner
}

func (p *Python) Install() {
	p.runner.Run([]string{
		"build",
		".",
		"--tag", p.image})
}

func (p *Python) Build() {
	log.Println("Building not supported in python")
}

func (p *Python) Test() {
	p.runner.Run([]string{
		"run",
		"-v", p.srcVolume,
		p.image,
		"nosetests", "test"})
}

func (p *Python) Lint() {
	p.runner.Run([]string{
		"run",
		"-v", p.srcVolume,
		p.image,
		"flake8", "setup.py", "test", p.name})
}

func (p *Python) Coverage() {
	p.runner.Run([]string{
		"run",
		"-v", p.srcVolume,
		p.image,
		"nosetests", "--with-coverage", "test"})
}

func (p *Python) CI() {
	p.Install()
	p.Lint()
	p.Coverage()
}

func (p *Python) Run(args []string) {
	p.runner.Run(append([]string{
		"run",
		"-v", p.srcVolume,
		p.image},
		args...))
}

func NewPython(root *Root, config *Config) *Python {
	srcVolume := fmt.Sprintf("%s:/usr/src/app", root.Root)
	runner := NewRunner(root, "docker")
	return &Python{config.Name, config.Name, srcVolume, runner}
}
