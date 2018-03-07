package meta_test

import (
	"github.com/stretchr/testify/mock"
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
	dotMeta *mocks.IDotMeta
	command *mocks.ICommand
	log     *mocks.ILog
}

func (suite *ParserTest) SetupTest() {
	suite.dotMeta = new(mocks.IDotMeta)
	suite.dotMeta.On("MoveToRoot").Return()
	suite.command = new(mocks.ICommand)
	suite.log = new(mocks.ILog)
	suite.parser = meta.NewParser(suite.dotMeta, suite.command, suite.log)
}

func (suite *ParserTest) TestInstall() {
	suite.command.On("Install").Return()
	suite.shell("meta install")
}

func (suite *ParserTest) TestBuild() {
	suite.command.On("Stage", "build", false).Return()
	suite.shell("meta build")
	suite.command.AssertExpectations(suite.T())

	suite.command.On("Stage", "build", true).Return()
	suite.shell("meta --image-only build")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestTest() {
	suite.command.On("Stage", "test", false).Return()
	suite.shell("meta test")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestLint() {
	suite.command.On("Stage", "lint", false).Return()
	suite.shell("meta lint")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestCoverage() {
	suite.command.On("Stage", "coverage", false).Return()
	suite.shell("meta coverage")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestEnter() {
	suite.command.On("Enter").Return()
	suite.shell("meta enter")
}

func (suite *ParserTest) TestCI() {
	suite.command.On("CI").Return()
	suite.shell("meta ci")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestRun() {
	suite.command.On("Run", []string{"command"}, false).Return()
	suite.shell("meta run -c \"command\"")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestVerify() {
	suite.command.On("Verify").Return()
	suite.shell("meta verify")
	suite.command.AssertExpectations(suite.T())
}

// func (suite *ParserTest) TestCreatePython() {
// 	suite.command.On("Create", "python", "name").Return()
// 	suite.shell("meta create python --mame name")
// 	suite.command.AssertExpectations(suite.T())
// }

func (suite *ParserTest) TestCreateGolang() {
	suite.command.On("Create", "golang", "name").Return()
	suite.shell("meta create golang --name name")
	suite.command.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestUnvalidCommand() {
	suite.log.On("Fatal", mock.Anything).Run(func(args mock.Arguments) {
		suite.Equal(args.Error(0).Error(), "unknown argument: unvalid")
	})
	suite.shell("meta unvalid")
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