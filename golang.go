package main

import (
	"fmt"
)

type Golang struct {
	name      string
	image     string
	srcVolume string
	runner    *Runner
}

func (g *Golang) Install() {
	g.runner.Run([]string{
		"build",
		".",
		"--tag", g.image})
}

func (g *Golang) Build() {
	g.runner.Run([]string{
		"run",
		"-v", g.srcVolume,
		g.image,
		"go", "build"})
}

func (g *Golang) Test() {
	g.runner.Run([]string{
		"run",
		"-v", g.srcVolume,
		g.image,
		"go", "test"})
}

func (g *Golang) Lint() {
	g.runner.Run([]string{
		"run",
		"-v", g.srcVolume,
		g.image,
		"go", "vet"})

	g.runner.Run([]string{
		"run",
		"-v", g.srcVolume,
		g.image,
		"go", "fmt"})
}

func (g *Golang) Coverage() {
	g.runner.Run([]string{
		"run",
		"-v", g.srcVolume,
		g.image,
		"go", "test", "-cover"})
}

func (g *Golang) CI() {
	g.Install()
	g.Build()
	g.Lint()
	g.Coverage()
}

func (g *Golang) Run(args []string) {
	g.runner.Run(append([]string{
		"run",
		"-v", g.srcVolume,
		g.image},
		args...))
}

func NewGolang(root *Root, config *Config) *Golang {
	srcVolume := fmt.Sprintf("%s:/go/src/%s", root.Root, config.Name)
	runner := NewRunner(root, "docker")
	return &Golang{config.Name, config.Name, srcVolume, runner}
}
