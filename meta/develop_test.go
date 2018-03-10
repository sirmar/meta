package meta_test

import (
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"testing"
)

type DevelopTest struct {
	suite.Suite
	util     *mocks.IUtil
	runner   *mocks.IRunner
	dotMeta  *mocks.IDotMeta
	settings *mocks.ISettings
	develop  meta.IDevelop
}

func (suite *DevelopTest) SetupTest() {
	suite.util = new(mocks.IUtil)
	suite.runner = new(mocks.IRunner)
	suite.dotMeta = new(mocks.IDotMeta)
	suite.settings = new(mocks.ISettings)
	suite.develop = meta.NewDevelop(suite.util, suite.runner, suite.dotMeta, suite.settings)
}

func (suite *DevelopTest) TestInstall() {
	suite.dotMeta.On("ReadMetaYml").Return(&meta.MetaYml{"name", "language"})
	suite.runner.On("Run", "docker", contains("build . --tag name")).Return()
	suite.develop.Install()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestEnter() {
	suite.dotMeta.On("ReadMetaYml").Return(&meta.MetaYml{"name", "language"})
	suite.runner.On("Run", "docker", contains("run -it name sh")).Return()
	suite.develop.Enter()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestStage() {
	suite.dotMeta.On("ReadMetaYml").Return(&meta.MetaYml{"name", "golang"})
	languageYml := map[string][]string{"build": []string{"cmd1", "cmd2"}}
	suite.settings.On("ReadLanguageYml", "golang").Return(&meta.LanguageYml{languageYml})

	suite.runner.On("Run", "docker", contains("run name cmd1")).Return()
	suite.runner.On("Run", "docker", contains("run name cmd2")).Return()
	suite.develop.Stage("build", true)
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestUpload() {
	suite.dotMeta.On("ReadMetaYml").Return(&meta.MetaYml{"name", "language"})
	suite.runner.On("Run", "docker", contains("push name")).Return()
	suite.develop.Upload()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestRunInGolang() {
	suite.dotMeta.On("ReadMetaYml").Return(&meta.MetaYml{"name", "golang"})
	suite.util.On("GetCwd").Return("/root")
	suite.runner.On("Run", "docker", contains("run -v /root:/go/src/name name go build")).Return()
	suite.develop.Run([]string{"go", "build"}, false)
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestRunInOtherLanguages() {
	suite.dotMeta.On("ReadMetaYml").Return(&meta.MetaYml{"name", "python"})
	suite.util.On("GetCwd").Return("/root")
	suite.runner.On("Run", "docker", contains("run -v /root:/usr/src/name name cmd")).Return()
	suite.develop.Run([]string{"cmd"}, false)
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestRunWithImageOnly() {
	suite.dotMeta.On("ReadMetaYml").Return(&meta.MetaYml{"name", "python"})
	suite.runner.On("Run", "docker", contains("run name cmd")).Return()
	suite.develop.Run([]string{"cmd"}, true)
	suite.runner.AssertExpectations(suite.T())
}

func TestDevelopSuite(t *testing.T) {
	suite.Run(t, new(DevelopTest))
}
