package main

import (
	"fmt"
)

type Golang struct {
	name string
	image string
	workDir string
	cacheVolume string
	srcVolume string
	runner *Runner
}

func (g* Golang) Install() {
	g.runner.Run([]string{
		"build",
		".",
		"--tag", g.image})

	g.runner.Run([]string{
		"run",
		"-w", g.workDir,
		"-v", g.cacheVolume,
		"-v", g.srcVolume,
		g.image,
		"go", "get", "-v", "-d"})
}

func (g* Golang) Build() {
	g.runner.Run([]string{
		"run",
		"-w", g.workDir,
		"-v", g.cacheVolume,
		"-v", g.srcVolume,
		g.image,
		"go", "build"})
}

func (g* Golang) Test() {
	g.runner.Run([]string{
		"run",
		"-w", g.workDir,
		"-v", g.cacheVolume,
		"-v", g.srcVolume,
		g.image,
		"go", "test"})
}

func (g* Golang) Lint() {
}

func (g* Golang) Coverage() {
}

func (g* Golang) CI() {
}

func (g* Golang) Run() {
	g.runner.Run([]string{
		"run",
		"-w", g.workDir,
		"-v", g.srcVolume,
		g.image,
		"meta"})
}

func NewGolang(root *Root, config *Config) *Golang {
	workDir := fmt.Sprintf("/go/src/%s", config.Name)
	cacheVolume := fmt.Sprintf("%s/.cache:/go", root.Root)
	srcVolume := fmt.Sprintf("%s:/go/src/%s", root.Root, config.Name)
	runner := NewRunner(root, "docker")
	return &Golang{config.Name, config.Name, workDir, cacheVolume, srcVolume, runner}
}
