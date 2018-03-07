package meta

//go:generate mockery -name=ICommand
type ICommand interface {
	Install()
	Enter()
	Stage(stage string, imageOnly bool)
	CI()
	Run(args []string, imageOnly bool)
	Create(language, name string)
	Verify()
}

type Command struct {
	develop IDevelop
	create  ICreate
	verify  IVerify
}

func NewCommand(develop IDevelop, create ICreate, verify IVerify) ICommand {
	return &Command{develop, create, verify}
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
