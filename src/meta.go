package main

import (
	"flag"
	"log"
	"os"
)

type Language interface {
	Install()
	Build()
	Test()
	Lint()
	Coverage()
}

func main() {
	install := flag.NewFlagSet("install", flag.ExitOnError)
	build := flag.NewFlagSet("build", flag.ExitOnError)
	test := flag.NewFlagSet("test", flag.ExitOnError)
	coverage := flag.NewFlagSet("coverage", flag.ExitOnError)
	lint := flag.NewFlagSet("lint", flag.ExitOnError)

	if len(os.Args) < 2 {
		log.Fatal("Usage: meta command")
	}

	switch os.Args[1] {
	case "install":
		install.Parse(os.Args[2:])
	case "build":
		build.Parse(os.Args[2:])
	case "test":
		test.Parse(os.Args[2:])
	case "coverage":
		coverage.Parse(os.Args[2:])
	case "lint":
		lint.Parse(os.Args[2:])
	default:
		log.Fatal("%q is not valid command.", os.Args[1])
	}

	var language Language
	root := NewRoot()
	config := NewConfig()

	switch config.Language {
	case "python":
		language = NewPython(root, config)
	}

	if install.Parsed() {
		language.Install()
	}
	if build.Parsed() {
		language.Build()
	}
	if test.Parsed() {
		language.Test()
	}
	if lint.Parsed() {
		language.Lint()
	}
	if coverage.Parsed() {
		language.Coverage()
	}
}
