package meta_test

import (
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"testing"
)

type CommandTest struct {
	suite.Suite
	command  meta.ICommand
	runner   *mocks.IRunner
	template *mocks.ITemplate
}

func (suite *CommandTest) SetupTest() {
	suite.runner = new(mocks.IRunner)
	suite.template = new(mocks.ITemplate)
	suite.command = meta.NewCommand(suite.runner, mockDotMeta(), suite.template)
}

func (suite *CommandTest) TestInstall() {
	suite.runner.On("Run", "docker", contains("build . --tag")).Return()
	suite.command.Install()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *CommandTest) TestEnter() {
	suite.runner.On("Run", "docker", contains("sh")).Return()
	suite.command.Enter()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *CommandTest) TestRun() {
	suite.runner.On("Run", "docker", contains("run this")).Return()
	suite.command.Run([]string{"run", "this"}, false)
	suite.runner.AssertExpectations(suite.T())
}

func (suite *CommandTest) TestCreate() {
	suite.template.On("Create", "name", "language").Return()
	suite.command.Create("name", "language")
	suite.runner.AssertExpectations(suite.T())
}

func (suite *CommandTest) TestLanguage() {
	suite.runner.On("Run", "docker", contains("run name cmd1")).Return()
	suite.runner.On("Run", "docker", contains("run name cmd2")).Return()
	suite.command.Language([]string{"cmd1", "cmd2"}, true)
	suite.runner.AssertExpectations(suite.T())

	suite.runner.On("Run", "docker", contains("run -v /root:/usr/src/name name cmd")).Return()
	suite.command.Language([]string{"cmd"}, false)
	suite.runner.AssertExpectations(suite.T())
}

func TestCommandSuite(t *testing.T) {
	suite.Run(t, new(CommandTest))
}
