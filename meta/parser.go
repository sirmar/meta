package meta

import (
	"github.com/devfacet/gocmd"
	"strings"
)

type Parser struct {
	languageYml *LanguageYml
	command     ICommand
	log         ILog
}

func NewParser(languageYml *LanguageYml, command ICommand, log ILog) *Parser {
	return &Parser{languageYml, command, log}
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
		Verify struct {
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
		p.log.Fatal(err)
	}

	if cmd.FlagArgs("Install") != nil {
		p.command.Install()
	} else if cmd.FlagArgs("Enter") != nil {
		p.command.Enter()
	} else if cmd.FlagArgs("Build") != nil {
		p.command.Language(p.languageYml.Build, flags.ImageOnly)
	} else if cmd.FlagArgs("Test") != nil {
		p.command.Language(p.languageYml.Test, flags.ImageOnly)
	} else if cmd.FlagArgs("Coverage") != nil {
		p.command.Language(p.languageYml.Coverage, flags.ImageOnly)
	} else if cmd.FlagArgs("Lint") != nil {
		p.command.Language(p.languageYml.Lint, flags.ImageOnly)
	} else if cmd.FlagArgs("CI") != nil {
		p.command.CI(p.languageYml.CI)
	} else if cmd.FlagArgs("Run") != nil {
		p.command.Run(strings.Split(flags.Run.Command, " "), flags.ImageOnly)
	} else if cmd.FlagArgs("Create") != nil {
		if cmd.FlagArgs("Create.Python") != nil {
			p.command.Create(flags.Create.Python.Name, "python")
		} else {
			p.log.Fatal("missing language for create command")
		}
	}
}
