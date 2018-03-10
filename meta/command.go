package meta

//go:generate mockery -name=ICommand
type ICommand interface {
	Setup()
	Install()
	Enter()
	Stage(stage string, imageOnly bool)
	Upload()
	CI()
	Run(args []string, imageOnly bool)
	Create(language, name string)
	Verify()
}

type Command struct {
	develop IDevelop
	create  ICreate
	verify  IVerify
	setup   ISetup
}

func NewCommand(develop IDevelop, create ICreate, verify IVerify, setup ISetup) ICommand {
	return &Command{develop, create, verify, setup}
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
