package meta

import (
	"fmt"
)

type Golang struct {
	Language
}

func NewGolang(runner IRunner, meta *DotMeta) *Golang {
	name := meta.Meta.Name
	srcVolume := fmt.Sprintf("%s:/go/src/%s", meta.Root, name)
	return &Golang{Language{name, name, srcVolume, runner}}
}

func (g *Golang) Build() {
	g.dockerRun("go", "build", "./...")
}

func (g *Golang) Test() {
	g.dockerRun("go", "test", "./...")
}

func (g *Golang) Lint() {
	g.dockerRun("go", "vet", "./...")
	g.dockerRun("go", "fmt", "./...")
}

func (g *Golang) Coverage() {
	g.dockerRun("go", "test", "-cover", "./...")
}

func (g *Golang) CI() {
	g.Install()
	g.Build()
	g.Lint()
	g.Coverage()
}
