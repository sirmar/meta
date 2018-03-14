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
	log     *mocks.ILog
	develop *mocks.IDevelop
	create  *mocks.ICreate
	verify  *mocks.IVerify
	setup   *mocks.ISetup
	release *mocks.IRelease
}

func (suite *ParserTest) SetupTest() {
	suite.dotMeta = new(mocks.IDotMeta)
	suite.dotMeta.On("MoveToRoot")
	suite.log = new(mocks.ILog)
	suite.develop = new(mocks.IDevelop)
	suite.create = new(mocks.ICreate)
	suite.verify = new(mocks.IVerify)
	suite.setup = new(mocks.ISetup)
	suite.release = new(mocks.IRelease)
	suite.parser = meta.NewParser(suite.dotMeta, suite.log, suite.develop, suite.create, suite.verify, suite.setup, suite.release)
}

func (suite *ParserTest) TestInstall() {
	suite.develop.On("Install", false)
	suite.shell("meta install")
}

func (suite *ParserTest) TestInstallNoCache() {
	suite.develop.On("Install", true)
	suite.shell("meta install --no-cache")
}

func (suite *ParserTest) TestBuild() {
	suite.develop.On("Stage", "build", false)
	suite.shell("meta build")
	suite.develop.AssertExpectations(suite.T())

	suite.develop.On("Stage", "build", true)
	suite.shell("meta --image-only build")
	suite.develop.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestTest() {
	suite.develop.On("Stage", "test", false)
	suite.shell("meta test")
	suite.develop.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestLint() {
	suite.develop.On("Stage", "lint", false)
	suite.shell("meta lint")
	suite.develop.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestCoverage() {
	suite.develop.On("Stage", "coverage", false)
	suite.shell("meta coverage")
	suite.develop.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestEnter() {
	suite.develop.On("Enter")
	suite.shell("meta enter")
}

func (suite *ParserTest) TestCI() {
	suite.develop.On("Install", false)
	suite.develop.On("Stage", "ci", true)
	suite.shell("meta ci")
	suite.develop.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestCINoCache() {
	suite.develop.On("Install", true)
	suite.develop.On("Stage", "ci", true)
	suite.shell("meta ci -n")
	suite.develop.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestRun() {
	suite.develop.On("Run", []string{"command"}, false)
	suite.shell("meta run -c \"command\"")
	suite.develop.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestVerify() {
	suite.verify.On("Files")
	suite.shell("meta verify")
	suite.verify.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestUpload() {
	suite.develop.On("Upload")
	suite.shell("meta upload")
	suite.develop.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestCreatePython() {
	suite.create.On("Template", "python", "name")
	suite.shell("meta create python --name name")
	suite.create.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestCreateGolang() {
	suite.create.On("Template", "golang", "name")
	suite.shell("meta create golang --name name")
	suite.create.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestCreateGeneral() {
	suite.create.On("Template", "general", "name")
	suite.shell("meta create general --name name")
	suite.create.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestSetup() {
	suite.setup.On("Run")
	suite.shell("meta setup")
	suite.setup.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestLogin() {
	suite.develop.On("Login")
	suite.shell("meta login")
	suite.develop.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestUnvalidCreate() {
	suite.log.On("Fatal", mock.Anything).Run(func(args mock.Arguments) {
		suite.Equal(args.String(0), "need language and name")
	})
	suite.shell("meta create")
}

func (suite *ParserTest) TestRelease() {
	suite.release.On("Create", "minor", "message")
	suite.shell("meta release minor -m message")
	suite.release.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestReleases() {
	suite.release.On("List")
	suite.shell("meta releases")
	suite.release.AssertExpectations(suite.T())
}

func (suite *ParserTest) TestDiff() {
	suite.release.On("Unreleased")
	suite.shell("meta diff")
	suite.release.AssertExpectations(suite.T())
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
