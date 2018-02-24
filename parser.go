package main

import (
	"github.com/akamensky/argparse"
	"log"
	"os"
	"strings"
)

func NewParser(language Language) {
	parser := argparse.NewParser("commands", "Simple example of argparse commands")

	install := parser.NewCommand("install", "install help")
	build := parser.NewCommand("build", "build help")
	test := parser.NewCommand("test", "test help")
	coverage := parser.NewCommand("coverage", "coverage help")
	lint := parser.NewCommand("lint", "lint help")
	ci := parser.NewCommand("ci", "ci help")

	run := parser.NewCommand("run", "run help")
	runCmd := run.String("c", "command", &argparse.Options{Required: true, Help: "Command to run inside docker container"})

	create := parser.NewCommand("create", "create help")
	python := create.NewCommand("python", "python help")

	validate := parser.NewCommand("validate", "validate help")

	err := parser.Parse(os.Args)

	if err != nil {
		log.Fatal(parser.Usage(err))
	}

	if install.Happened() {
		language.Install()
	} else if build.Happened() {
		language.Build()
	} else if test.Happened() {
		language.Test()
	} else if lint.Happened() {
		language.Lint()
	} else if coverage.Happened() {
		language.Coverage()
	} else if ci.Happened() {
		language.CI()
	} else if run.Happened() {
		language.Run(strings.Split(*runCmd, " "))
	} else if create.Happened() {
		if python.Happened() {
			log.Println("create python")
		}
	} else if validate.Happened() {
	}
}
