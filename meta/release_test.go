package meta_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"testing"
)

type ReleaseTest struct {
	suite.Suite
	util    *mocks.IUtil
	runner  *mocks.IRunner
	log     *mocks.ILog
	release meta.IRelease
}

func (suite *ReleaseTest) SetupTest() {
	suite.util = new(mocks.IUtil)
	suite.runner = new(mocks.IRunner)
	suite.log = new(mocks.ILog)
	suite.release = meta.NewRelease(suite.util, suite.runner, suite.log)
}

func (suite *ReleaseTest) TestCreateFromNoCurrentRelease() {
	suite.givenLatestRelease("")
	suite.givenReleaseConfirmed(true)
	suite.runner.On("Run", "git", contains("tag 0.0.1 -a -m message"))
	suite.runner.On("Run", "git", contains("push origin 0.0.1"))
	suite.release.Create("patch", "message")
	suite.runner.AssertExpectations(suite.T())
}

func (suite *ReleaseTest) TestCreatePatch() {
	suite.givenLatestRelease("1.2.3")
	suite.givenReleaseConfirmed(true)
	suite.runner.On("Run", "git", contains("tag 1.2.4 -a -m message"))
	suite.runner.On("Run", "git", contains("push origin 1.2.4"))
	suite.release.Create("patch", "message")
	suite.runner.AssertExpectations(suite.T())
}

func (suite *ReleaseTest) TestCreateMinor() {
	suite.givenLatestRelease("1.2.3")
	suite.givenReleaseConfirmed(true)
	suite.runner.On("Run", "git", contains("tag 1.3.0 -a -m message"))
	suite.runner.On("Run", "git", contains("push origin 1.3.0"))
	suite.release.Create("minor", "message")
	suite.runner.AssertExpectations(suite.T())
}

func (suite *ReleaseTest) TestCreateMajor() {
	suite.givenLatestRelease("1.2.3")
	suite.givenReleaseConfirmed(true)
	suite.runner.On("Run", "git", contains("tag 2.0.0 -a -m message"))
	suite.runner.On("Run", "git", contains("push origin 2.0.0"))
	suite.release.Create("major", "message")
	suite.runner.AssertExpectations(suite.T())
}

func (suite *ReleaseTest) TestCreateButNotConfirmed() {
	suite.givenLatestRelease("1.2.3")
	suite.givenReleaseConfirmed(false)
	suite.release.Create("major", "message")
	suite.runner.AssertExpectations(suite.T())
}

func (suite *ReleaseTest) TestList() {
	suite.runner.On("Run", "git", contains("tag -n --sort version:refname -l [0-9]*"))
	suite.release.List()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *ReleaseTest) TestUnreleasedFromNoCurrentRelease() {
	suite.givenLatestRelease("")
	suite.release.Unreleased()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *ReleaseTest) TestUnreleased() {
	suite.givenLatestRelease("1.2.3")
	suite.runner.On("Run", "git", contains("log 1.2.3 ..HEAD --oneline"))
	suite.release.Unreleased()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *ReleaseTest) givenLatestRelease(release string) {
	suite.runner.On("Output", "git", contains("tag --sort version:refname -l [0-9]*")).Return(release)
}

func (suite *ReleaseTest) givenReleaseConfirmed(yesno bool) {
	suite.util.On("YesNo", mock.Anything).Return(yesno)
}

func TestReleaseSuite(t *testing.T) {
	suite.Run(t, new(ReleaseTest))
}
