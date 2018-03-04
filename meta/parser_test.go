package meta_test

import (
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"os"
	"strings"
	"testing"
)

type ParserTest struct {
	suite.Suite
	parser  *meta.Parser
	command *mocks.ICommand
}

func (suite *ParserTest) SetupTest() {
	suite.command = new(mocks.ICommand)
	suite.parser = meta.NewParser(mockLanguageYml(), suite.command, new(mocks.ILog))
}

func (suite *ParserTest) TestInstall() {
	suite.command.On("Install").Return()
	suite.shell("meta install")
}

func (suite *ParserTest) TestBuild() {
	suite.command.On("Language", []string{"build"}, false).Return()
	suite.shell("meta build")
	suite.command.AssertExpectations(suite.T())

	suite.command.On("Language", []string{"build"}, true).Return()
	suite.shell("meta --image-only build")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestTest() {
	suite.command.On("Language", []string{"test"}, false).Return()
	suite.shell("meta test")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestLint() {
	suite.command.On("Language", []string{"lint"}, false).Return()
	suite.shell("meta lint")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestCoverage() {
	suite.command.On("Language", []string{"coverage"}, false).Return()
	suite.shell("meta coverage")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestEnter() {
	suite.command.On("Enter").Return()
	suite.shell("meta enter")
}

func (suite *ParserTest) TestCI() {
	suite.command.On("Language", []string{"ci1", "ci2"}, false).Return()
	suite.shell("meta ci")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestRun() {
	suite.command.On("Run", []string{"command"}, false).Return()
	suite.shell("meta run -c \"command\"")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) shell(cmd string) {
	oldArgs := os.Args
	os.Args = strings.Split(cmd, " ")
	suite.parser.Run()
	os.Args = oldArgs
}

func TestParserSuite(t *testing.T) {
	suite.Run(t, new(ParserTest))
}
