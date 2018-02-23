package main

import (
	"fmt"
)

type Golang struct {
	name string
	image string
	cacheVolume string
	srcVolume string
	runner *Runner
}

func (g* Golang) Install() {
	g.runner.Run([]string{"build", ".", "--tag", g.image})
	g.runner.Run([]string{
		"run",
		"-w", "/go/src/app",
		"-v", g.cacheVolume,
		"-v", g.srcVolume,
		"-e", "GOOS=darwin",
		"-e", "GOARCH=386",
		g.image,
		"go", "get", "-v", "-d"})
}

func (g* Golang) Build() {
	g.runner.Run([]string{
		"run",
		"-w", "/go/src/app",
		"-v", g.cacheVolume,
		"-v", g.srcVolume,
		"-e", "GOOS=darwin",
		"-e", "GOARCH=386",
		g.image,
		"go", "build"})
}

func (g* Golang) Test() {
}

func (g* Golang) Lint() {
}

func (g* Golang) Coverage() {
}

func (g* Golang) CI() {
}

func NewGolang(root *Root, config *Config) *Golang {
	cacheVolume := fmt.Sprintf("%s/.cache:/go", root.Root)
	srcVolume := fmt.Sprintf("%s:/go/src/app", root.Root)
	runner := NewRunner(root, "docker")
	return &Golang{config.Name, config.Name, cacheVolume, srcVolume, runner}
}
