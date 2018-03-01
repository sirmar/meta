package main

import (
	"log"
)

type ILanguage interface {
	Install()
	Build()
	Test()
	Lint()
	Coverage()
	CI()
	Enter()
	Run(args []string)
	SetImageOnly()
}

type Language struct {
	name      string
	image     string
	srcVolume string
	runner    IRunner
}

func NewLanguage(runner IRunner, meta *Meta) ILanguage {
	switch meta.Meta.Language {
	case "python":
		return NewPython(runner, meta)
	case "golang":
		return NewGolang(runner, meta)
	default:
		log.Fatal(meta.Meta.Language, "not supported!")
		return NewPython(runner, meta)
	}
}

func (l *Language) Install() {
	l.dockerBuild()
}

func (l *Language) Build() {
	log.Println("Building not supported in this language ")
}

func (l *Language) Enter() {
	if l.useImageOnly() {
		l.runner.Run([]string{"run", "-it", l.image, "sh"})
	} else {
		l.runner.Run([]string{"run", "-it", "-v", l.srcVolume, l.image, "sh"})
	}
}

func (l *Language) Run(args []string) {
	l.dockerRun(args...)
}

func (l *Language) SetImageOnly() {
	l.srcVolume = ""
}

func (l *Language) dockerRun(cmds ...string) {
	if l.useImageOnly() {
		l.runner.Run(append([]string{"run", l.image}, cmds...))
	} else {
		l.runner.Run(append([]string{"run", "-v", l.srcVolume, l.image}, cmds...))
	}
}

func (l *Language) dockerBuild() {
	l.runner.Run([]string{"build", ".", "--tag", l.image})
}

func (l *Language) useImageOnly() bool {
	return l.srcVolume == ""
}
