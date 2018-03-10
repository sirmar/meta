package meta_test

import (
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"testing"
)

type CommandTest struct {
	suite.Suite
	command meta.ICommand
	develop *mocks.IDevelop
	create  *mocks.ICreate
	verify  *mocks.IVerify
	setup   *mocks.ISetup
}

func (suite *CommandTest) SetupTest() {
	suite.develop = new(mocks.IDevelop)
	suite.create = new(mocks.ICreate)
	suite.verify = new(mocks.IVerify)
	suite.setup = new(mocks.ISetup)
	suite.command = meta.NewCommand(suite.develop, suite.create, suite.verify, suite.setup)
}

func (suite *CommandTest) TestSetup() {
	suite.setup.On("Run").Return()
	suite.command.Setup()
	suite.setup.AssertExpectations(suite.T())
}

func (suite *CommandTest) TestInstall() {
	suite.develop.On("Install").Return()
	suite.command.Install()
	suite.develop.AssertExpectations(suite.T())
}

func (suite *CommandTest) TestEnter() {
	suite.develop.On("Enter").Return()
	suite.command.Enter()
	suite.develop.AssertExpectations(suite.T())
}

func (suite *CommandTest) TestState() {
	suite.develop.On("Stage", "build", false).Return()
	suite.command.Stage("build", false)
	suite.develop.AssertExpectations(suite.T())
}

func (suite *CommandTest) TestCI() {
	suite.develop.On("Install").Return()
	suite.develop.On("Stage", "ci", true).Return()
	suite.command.CI()
	suite.develop.AssertExpectations(suite.T())
}

func (suite *CommandTest) TestRun() {
	suite.develop.On("Run", []string{"args"}, false).Return()
	suite.command.Run([]string{"args"}, false)
	suite.develop.AssertExpectations(suite.T())
}

func (suite *CommandTest) TestCreate() {
	suite.create.On("Template", "language", "name").Return()
	suite.command.Create("language", "name")
	suite.create.AssertExpectations(suite.T())
}

func (suite *CommandTest) TestVerify() {
	suite.verify.On("Files").Return()
	suite.command.Verify()
	suite.verify.AssertExpectations(suite.T())
}

func TestCommandSuite(t *testing.T) {
	suite.Run(t, new(CommandTest))
}
