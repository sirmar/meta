package main

import (
	"fmt"
)

type Python struct {
	pkgname string
	image string
	volume string
	runner *Runner
}

func (p* Python) Install() {
	p.runner.Run([]string{"build", ".", "--tag", p.image})
}

func (p* Python) Build() {
	fmt.Println("Building not supported in python")
}

func (p* Python) Test() {
	p.runner.Run([]string{"run", "-v", p.volume, p.image, "nosetests", "test"})
}

func (p* Python) Lint() {
	p.runner.Run([]string{"run", "-v", p.volume, p.image, "flake8", "setup.py", "test", p.pkgname})
}

func (p* Python) Coverage() {
	p.runner.Run([]string{"run", "-v", p.volume, p.image, "nosetests", "--with-coverage", "test"})
}

func NewPython(root *Root, config *Config) *Python {
	volume := fmt.Sprintf("%s:/usr/src/app", root.Root)
	runner := NewRunner(root, "docker")
	return &Python{config.Name, config.Name, volume, runner}
}
