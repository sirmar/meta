package meta

import (
	"github.com/devfacet/gocmd"
	"strings"
)

type CommandLine struct {
	Help      bool     `short:"h" long:"help" description:"Display usage" global:"true"`
	ImageOnly bool     `short:"i" long:"image-only" description:"Run based on docker image without using any volumes to see local changes" global:"true"`
	Version   bool     `short:"v" long:"version" description:"Display version"`
	Install   struct{} `command:"install" description:"Build the docker image that will install all dependencies"`
	Build     struct{} `command:"build" description:"Build the current project"`
	Test      struct{} `command:"test" description:"Run tests on the current project"`
	Coverage  struct{} `command:"coverage" description:"Run coverage on the current project"`
	Lint      struct{} `command:"lint" description:"Run lint tools on the current project"`
	CI        struct{} `command:"ci" description:"Install, build, test, coverage and lint in one command"`
	Enter     struct{} `command:"enter" description:"Enter docker image"`
	Run       struct {
		Command string `short:"c" long:"command" required:"true" description:"Command"`
	} `command:"run" description:"run description"`
	Create struct {
		// Python struct {
		// 	Name string `short:"m" long:"mame" required:"true" description:"Name of new project"`
		// } `command:"python" description:"Create a template python project"`
		Golang struct {
			Name string `short:"n" long:"name" required:"true" description:"Name of new project"`
		} `command:"golang" description:"Create a template golang project"`
	} `command:"create"`
	Verify struct{} `command:"verify" description:"Verify then project is suitable for this tool"`
}

type Parser struct {
	dotMeta IDotMeta
	command ICommand
	log     ILog
}

func NewParser(dotMeta IDotMeta, command ICommand, log ILog) *Parser {
	return &Parser{dotMeta, command, log}
}

func (self *Parser) Run() {
	flags := new(CommandLine)

	cmd, err := gocmd.New(gocmd.Options{
		Name:        "meta",
		Version:     "0.0.1",
		Description: "Unified interface for development",
		Flags:       flags,
		AutoHelp:    true,
		AutoVersion: true,
		AnyError:    true,
	})

	if err != nil {
		self.log.Fatal(err)
		return
	}

	if self.dotMetaIsRequired(cmd) {
		self.dotMeta.MoveToRoot()
		self.metaCommands(cmd, flags)
	} else {
		self.otherCommands(cmd, flags)
	}
}

func (self *Parser) metaCommands(cmd *gocmd.Cmd, flags *CommandLine) {
	if cmd.FlagArgs("Install") != nil {
		self.command.Install()
	} else if cmd.FlagArgs("Enter") != nil {
		self.command.Enter()
	} else if cmd.FlagArgs("Build") != nil {
		self.command.Stage("build", flags.ImageOnly)
	} else if cmd.FlagArgs("Test") != nil {
		self.command.Stage("test", flags.ImageOnly)
	} else if cmd.FlagArgs("Coverage") != nil {
		self.command.Stage("coverage", flags.ImageOnly)
	} else if cmd.FlagArgs("Lint") != nil {
		self.command.Stage("lint", flags.ImageOnly)
	} else if cmd.FlagArgs("CI") != nil {
		self.command.CI()
	} else if cmd.FlagArgs("Run") != nil {
		self.command.Run(strings.Split(flags.Run.Command, " "), flags.ImageOnly)
	} else if cmd.FlagArgs("Verify") != nil {
		self.command.Verify()
	}
}

func (self *Parser) otherCommands(cmd *gocmd.Cmd, flags *CommandLine) {
	if cmd.FlagArgs("Create") != nil {
		if cmd.FlagArgs("Create.Python") != nil {
			//self.command.Create("python", flags.Create.Python.Name)
		} else if cmd.FlagArgs("Create.Golang") != nil {
			self.command.Create("golang", flags.Create.Golang.Name)
		}
	}
}

func (self *Parser) dotMetaIsRequired(cmd *gocmd.Cmd) bool {
	return cmd.FlagArgs("Create") == nil
}
