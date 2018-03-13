package meta

//go:generate mockery -name=ICommand
type ICommand interface {
	Setup()
	Install()
	Enter()
	Stage(stage string, imageOnly bool)
	Upload()
	Login()
	CI()
	Run(args []string, imageOnly bool)
	Create(language, name string)
	Verify()
	Release(level, message string)
	Releases()
	Diff()
}

type Command struct {
	develop IDevelop
	create  ICreate
	verify  IVerify
	setup   ISetup
	release IRelease
}

func NewCommand(develop IDevelop, create ICreate, verify IVerify, setup ISetup, release IRelease) ICommand {
	return &Command{develop, create, verify, setup, release}
}

func (self *Command) Install() {
	self.develop.Install()
}

func (self *Command) Enter() {
	self.develop.Enter()
}

func (self *Command) Stage(stage string, imageOnly bool) {
	self.develop.Stage(stage, imageOnly)
}

func (self *Command) Upload() {
	self.develop.Upload()
}

func (self *Command) CI() {
	self.Install()
	self.develop.Stage("ci", true)
}

func (self *Command) Run(args []string, imageOnly bool) {
	self.develop.Run(args, imageOnly)
}

func (self *Command) Create(language, name string) {
	self.create.Template(language, name)
}

func (self *Command) Verify() {
	self.verify.Files()
}

func (self *Command) Setup() {
	self.setup.Run()
}

func (self *Command) Login() {
	self.develop.Login()
}

func (self *Command) Release(level, message string) {
	self.release.Create(level, message)
}

func (self *Command) Releases() {
	self.release.List()
}

func (self *Command) Diff() {
	self.release.Unreleased()
}
