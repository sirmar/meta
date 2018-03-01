package main

import (
	"github.com/devfacet/gocmd"
	"log"
	"strings"
)

type Parser struct {
	language ILanguage
	template ITemplate
}

func NewParser(language ILanguage, template ITemplate) *Parser {
	return &Parser{language, template}
}

func (p *Parser) Run() {
	flags := struct {
		Help      bool     `short:"h" long:"help" description:"Display usage" global:"true"`
		ImageOnly bool     `short:"i" long:"image-only" description:"Display usage" global:"true"`
		Version   bool     `short:"v" long:"version" description:"Display version"`
		Install   struct{} `command:"install" description:"install description"`
		Build     struct{} `command:"build" description:"build description"`
		Test      struct{} `command:"test" description:"test description"`
		Coverage  struct{} `command:"coverage" description:"coverage description"`
		Lint      struct{} `command:"lint" description:"lint description"`
		CI        struct{} `command:"ci" description:"ci description"`
		Enter     struct{} `command:"enter" description:"enter description"`
		Run       struct {
			Command string `short:"c" long:"command" required:"true" description:"Command"`
		} `command:"run" description:"run description"`
		Create struct {
			Python struct {
				Name string `short:"n" long:"name" required:"true" description:"Name"`
			} `command:"python" description:"python description"`
		} `command:"create" description:"create description"`
		verify struct {
			Python struct{} `command:"python" description:"python description"`
		} `command:"verify" description:"verify description"`
	}{}

	cmd, err := gocmd.New(gocmd.Options{
		Name:        "meta",
		Version:     "0.0.1",
		Description: "Unified interface for development",
		Flags:       &flags,
		AutoHelp:    true,
		AutoVersion: true,
		AnyError:    true,
	})

	if err != nil {
		log.Fatal(err)
	}

	if flags.ImageOnly {
		p.language.SetImageOnly()
	}

	if cmd.FlagArgs("Install") != nil {
		p.language.Install()
	} else if cmd.FlagArgs("Build") != nil {
		p.language.Build()
	} else if cmd.FlagArgs("Test") != nil {
		p.language.Test()
	} else if cmd.FlagArgs("Coverage") != nil {
		p.language.Coverage()
	} else if cmd.FlagArgs("Lint") != nil {
		p.language.Lint()
	} else if cmd.FlagArgs("CI") != nil {
		p.language.CI()
	} else if cmd.FlagArgs("Enter") != nil {
		p.language.Enter()
	} else if cmd.FlagArgs("Run") != nil {
		p.language.Run(strings.Split(flags.Run.Command, " "))
	} else if cmd.FlagArgs("Create") != nil {
		if cmd.FlagArgs("Create.Python") != nil {
			p.template.Create(flags.Create.Python.Name, "python")
		}
	}
}
