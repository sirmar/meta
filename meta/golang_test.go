package meta_test

import (
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"testing"
)

type GolangTest struct {
	suite.Suite
	golang *meta.Golang
	runner *mocks.IRunner
}

func (suite *GolangTest) SetupTest() {
	suite.runner = new(mocks.IRunner)
	suite.golang = meta.NewGolang(
		suite.runner,
		&meta.DotMeta{&meta.MetaYml{"test", "golang"}, "/root", new(mocks.IUtil)})
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
