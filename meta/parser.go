package meta

import (
	"github.com/devfacet/gocmd"
	"strings"
)

type CommandLine struct {
	Help      bool     `short:"h" long:"help" description:"Display usage" global:"true"`
	ImageOnly bool     `short:"i" long:"image-only" description:"Run based on docker image without using any volumes to see local changes" global:"true"`
	Version   bool     `long:"version" description:"Display version"`
	Verbose   bool     `short:"v" description:"Verbose output"`
	Quiet   bool     `short:"q" description:"No logging"`
	Setup     struct{} `command:"setup" description:"Setup up your meta configuration"`
	Install   struct {
		NoCache bool `short:"n" long:"no-cache" description:"Build without cache"`
	} `command:"install" description:"Build the docker image that will install all dependencies"`
	Build    struct{} `command:"build" description:"Build the current project"`
	Test     struct{} `command:"test" description:"Run tests on the current project"`
	Coverage struct{} `command:"coverage" description:"Run coverage on the current project"`
	Lint     struct{} `command:"lint" description:"Run lint tools on the current project"`
	CI       struct {
		NoCache bool `short:"n" long:"no-cache" description:"Build without cache"`
	} `command:"ci" description:"Install, build, test, coverage and lint in one command"`
	Enter struct{} `command:"enter" description:"Enter docker image"`

	Run struct {
		Command string `short:"c" long:"command" required:"true" description:"Command"`
	} `command:"run" description:"run description"`

	Create struct {
		Python struct {
			Name string `short:"n" long:"name" required:"true" description:"Name of new project"`
		} `command:"python" description:"Create a template python project"`
		Golang struct {
			Name string `short:"n" long:"name" required:"true" description:"Name of new project"`
		} `command:"golang" description:"Create a template golang project"`
		General struct {
			Name string `short:"n" long:"name" required:"true" description:"Name of new project"`
		} `command:"general" description:"Create a minimal template project"`
	} `command:"create"`

	Verify struct{} `command:"verify" description:"Verify then project is suitable for this tool"`
	Upload struct{} `command:"upload" description:"Upload docker image/package to server"`
	Login  struct{} `command:"login" description:"Login to docker registry"`

	Release struct {
		Patch struct {
			Message string `short:"m" long:"message" required:"true" description:"Release message"`
		} `command:"patch" description:"Create patch release"`
		Minor struct {
			Message string `short:"m" long:"message" required:"true" description:"Release message"`
		} `command:"minor" description:"Create minor release"`
		Major struct {
			Message string `short:"m" long:"message" required:"true" description:"Release message"`
		} `command:"major" description:"Create major release"`
	} `command:"release"`
	Releases struct{} `command:"releases" description:"List all releases"`
	Diff     struct{} `command:"diff" description:"Show diff since last release"`
}

type Parser struct {
	dotMeta IDotMeta
	log     ILog
	develop IDevelop
	create  ICreate
	verify  IVerify
	setup   ISetup
	release IRelease
}

func NewParser(dotMeta IDotMeta, log ILog, develop IDevelop, create ICreate, verify IVerify, setup ISetup, release IRelease) *Parser {
	return &Parser{dotMeta, log, develop, create, verify, setup, release}
}

func (self *Parser) Run() {
	flags := new(CommandLine)

	cmd, err := gocmd.New(gocmd.Options{
		Name:        "meta",
		Version:     "0.0.1",
		Description: "Unified interface for development",
		Flags:       flags,
		AutoHelp:    true,
		AutoVersion: false,
		AnyError:    true,
	})

	if err != nil {
		self.log.Fatal(err)
	}

	if flags.Verbose {
		self.log.SetVerbose()
	} else if flags.Quiet {
		self.log.SetQuiet()
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
		self.develop.Install(flags.Install.NoCache)
	} else if cmd.FlagArgs("Enter") != nil {
		self.develop.Enter()
	} else if cmd.FlagArgs("Build") != nil {
		self.develop.Stage("build", flags.ImageOnly)
	} else if cmd.FlagArgs("Test") != nil {
		self.develop.Stage("test", flags.ImageOnly)
	} else if cmd.FlagArgs("Coverage") != nil {
		self.develop.Stage("coverage", flags.ImageOnly)
	} else if cmd.FlagArgs("Lint") != nil {
		self.develop.Stage("lint", flags.ImageOnly)
	} else if cmd.FlagArgs("CI") != nil {
		self.develop.Install(flags.CI.NoCache)
		self.develop.Stage("ci", true)
	} else if cmd.FlagArgs("Run") != nil {
		self.develop.Run(strings.Split(flags.Run.Command, " "), flags.ImageOnly)
	} else if cmd.FlagArgs("Verify") != nil {
		self.verify.Files()
	} else if cmd.FlagArgs("Upload") != nil {
		self.develop.Upload()
	}
}

func (self *Parser) otherCommands(cmd *gocmd.Cmd, flags *CommandLine) {
	if cmd.FlagArgs("Create") != nil {
		if cmd.FlagArgs("Create.Python") != nil {
			self.create.Template("python", flags.Create.Python.Name)
		} else if cmd.FlagArgs("Create.Golang") != nil {
			self.create.Template("golang", flags.Create.Golang.Name)
		} else if cmd.FlagArgs("Create.General") != nil {
			self.create.Template("general", flags.Create.General.Name)
		} else {
			self.log.Fatal("need language and name")
		}
	} else if cmd.FlagArgs("Setup") != nil {
		self.setup.Run()
	} else if cmd.FlagArgs("Login") != nil {
		self.develop.Login()
	} else if cmd.FlagArgs("Release") != nil {
		if cmd.FlagArgs("Release.Patch") != nil {
			self.release.Create("patch", flags.Release.Patch.Message)
		} else if cmd.FlagArgs("Release.Minor") != nil {
			self.release.Create("minor", flags.Release.Minor.Message)
		} else if cmd.FlagArgs("Release.Major") != nil {
			self.release.Create("major", flags.Release.Major.Message)
		} else {
			self.log.Fatal("need patch, minor or major")
		}
	} else if cmd.FlagArgs("Releases") != nil {
		self.release.List()
	} else if cmd.FlagArgs("Diff") != nil {
		self.release.Unreleased()
	}
}

func (self *Parser) dotMetaIsRequired(cmd *gocmd.Cmd) bool {
	return cmd.FlagArgs("Create") == nil &&
		cmd.FlagArgs("Setup") == nil &&
		cmd.FlagArgs("Login") == nil &&
		cmd.FlagArgs("Release") == nil &&
		cmd.FlagArgs("Releases") == nil &&
		cmd.FlagArgs("Diff") == nil
}
