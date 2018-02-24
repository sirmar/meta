package main

import (
	"meta/mocks"
	"os"
	"testing"
)

//go:generate mockery -name=Language

func TestBuildCommand(t *testing.T) {
	language := new(mocks.Language)
	language.On("Build").Return()
	meta("build", language)
}

func meta(cmd string, language Language) {
	oldArgs := os.Args
	os.Args = []string{"meta", cmd}
	NewParser(language)
	os.Args = oldArgs
}
