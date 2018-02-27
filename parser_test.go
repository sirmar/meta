package main_test

import (
	"github.com/stretchr/testify/suite"
	"meta"
	"meta/mocks"
	"os"
	"testing"
)

//go:generate mockery -name=ILanguage

type ParserTest struct {
	suite.Suite
	parser   *main.Parser
	language *mocks.ILanguage
}

func (suite *ParserTest) SetupTest() {
	suite.language = new(mocks.ILanguage)
	suite.parser = main.NewParser(suite.language)
}

func (suite *ParserTest) TestInstall() {
	suite.language.On("Install").Return()
	suite.meta("install")
}

func (suite *ParserTest) TestBuild() {
	suite.language.On("Build").Return()
	suite.meta("build")
}

func (suite *ParserTest) TestTest() {
	suite.language.On("Test").Return()
	suite.meta("test")
}

func (suite *ParserTest) TestLint() {
	suite.language.On("Lint").Return()
	suite.meta("lint")
}

func (suite *ParserTest) TestCoverage() {
	suite.language.On("Coverage").Return()
	suite.meta("coverage")
}

func (suite *ParserTest) TestEnter() {
	suite.language.On("Enter").Return()
	suite.meta("enter")
}

func (suite *ParserTest) TestCI() {
	suite.language.On("CI").Return()
	suite.meta("ci")
}

func (suite *ParserTest) TestRun() {
	// suite.language.On("run", []string{"go", "generate"}).Return()
	// suite.meta("run", "-c", "go generate")
}

func (suite *ParserTest) meta(cmd ...string) {
	oldArgs := os.Args
	os.Args = append([]string{"meta"}, cmd...)
	suite.parser.Run()
	os.Args = oldArgs
	suite.language.AssertExpectations(suite.T())
}

func TestParserSuite(t *testing.T) {
	suite.Run(t, new(ParserTest))
}
