package main

import (
	"fmt"
)

type Golang struct {
	Language
}

func NewGolang(runner IRunner, root *Root, config *Config) *Golang {
	srcVolume := fmt.Sprintf("%s:/go/src/%s", root.Root, config.Name)
	return &Golang{Language{config.Name, config.Name, srcVolume, runner}}
}

func (g *Golang) Build() {
	g.dockerRun("go", "build")
}

func (g *Golang) Test() {
	g.dockerRun("go", "test")
}

func (g *Golang) Lint() {
	g.dockerRun("go", "vet")
	g.dockerRun("go", "fmt")
}

func (g *Golang) Coverage() {
	g.dockerRun("go", "test", "-cover")
}

func (g *Golang) CI() {
	g.Install()
	g.Build()
	g.Lint()
	g.Coverage()
}
