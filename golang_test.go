package main_test

import (
	"github.com/stretchr/testify/suite"
	"meta"
	"meta/mocks"
	"testing"
)

//go:generate mockery -name=IRunner

type GolangTest struct {
	suite.Suite
	golang   *main.Golang
	runner  *mocks.IRunner
}

func (suite *GolangTest) SetupTest() {
	suite.runner = new(mocks.IRunner)
	suite.golang = main.NewGolang(
		suite.runner,
		&main.Root{"/root", "/root"},
		&main.Config{"test", "golang"})
}

func (suite *GolangTest) TestInstall() {
	suite.runner.On("Run", contains("build")).Return()
	suite.golang.Install()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *GolangTest) TestBuild() {
	suite.runner.On("Run", contains("go build")).Return()
	suite.golang.Build()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *GolangTest) TestTest() {
	suite.runner.On("Run", contains("go test")).Return()
	suite.golang.Test()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *GolangTest) TestLint() {
	suite.runner.On("Run", contains("go vet")).Return()
	suite.runner.On("Run", contains("go fmt")).Return()
	suite.golang.Lint()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *GolangTest) TestCoverage() {
	suite.runner.On("Run", contains("go test -cover")).Return()
	suite.golang.Coverage()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *GolangTest) TestEnter() {
	suite.runner.On("Run", contains("sh")).Return()
	suite.golang.Enter()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *GolangTest) TestRun() {
	suite.runner.On("Run", contains("run this")).Return()
	suite.golang.Run([]string{"run", "this"})
	suite.runner.AssertExpectations(suite.T())
}

func (suite *GolangTest) TestCI() {
	suite.runner.On("Run", contains("build")).Return()
	suite.runner.On("Run", contains("go build")).Return()
	suite.runner.On("Run", contains("go vet")).Return()
	suite.runner.On("Run", contains("go fmt")).Return()
	suite.runner.On("Run", contains("go test -cover")).Return()
	suite.golang.CI()
	suite.runner.AssertExpectations(suite.T())
}

func TestGolangSuite(t *testing.T) {
	suite.Run(t, new(GolangTest))
}
