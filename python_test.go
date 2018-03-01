package main_test

import (
	"github.com/stretchr/testify/suite"
	"meta"
	"meta/mocks"
	"testing"
)

//go:generate mockery -name=IRunner

type PythonTest struct {
	suite.Suite
	python *main.Python
	runner *mocks.IRunner
}

func (suite *PythonTest) SetupTest() {
	suite.runner = new(mocks.IRunner)
	suite.python = main.NewPython(
		suite.runner,
		&main.Meta{&main.MetaYml{"name", "python"}, "/root", "/root"})
}

func (suite *PythonTest) TestInstall() {
	suite.runner.On("Run", contains("build")).Return()
	suite.python.Install()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *PythonTest) TestTest() {
	suite.runner.On("Run", contains("nosetests test")).Return()
	suite.python.Test()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *PythonTest) TestLint() {
	suite.runner.On("Run", contains("flake8 setup.py test name")).Return()
	suite.python.Lint()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *PythonTest) TestCoverage() {
	suite.runner.On("Run", contains("nosetests --with-coverage test")).Return()
	suite.python.Coverage()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *PythonTest) TestEnter() {
	suite.runner.On("Run", contains("sh")).Return()
	suite.python.Enter()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *PythonTest) TestRun() {
	suite.runner.On("Run", contains("run this")).Return()
	suite.python.Run([]string{"run", "this"})
	suite.runner.AssertExpectations(suite.T())
}

func (suite *PythonTest) TestCI() {
	suite.runner.On("Run", contains("build")).Return()
	suite.runner.On("Run", contains("flake8 setup.py test name")).Return()
	suite.runner.On("Run", contains("nosetests --with-coverage test")).Return()
	suite.python.CI()
	suite.runner.AssertExpectations(suite.T())
}

func TestPythonSuite(t *testing.T) {
	suite.Run(t, new(PythonTest))
}
