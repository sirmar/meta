package meta_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"testing"
)

type DevelopTest struct {
	suite.Suite
	util     *mocks.IUtil
	log      *mocks.ILog
	runner   *mocks.IRunner
	dotMeta  *mocks.IDotMeta
	settings *mocks.ISettings
	template *mocks.ITemplate
	develop  meta.IDevelop
}

func (suite *DevelopTest) SetupTest() {
	suite.util = new(mocks.IUtil)
	suite.log = new(mocks.ILog)
	suite.runner = new(mocks.IRunner)
	suite.dotMeta = new(mocks.IDotMeta)
	suite.settings = new(mocks.ISettings)
	suite.template = new(mocks.ITemplate)
	suite.develop = meta.NewDevelop(suite.util, suite.log, suite.runner, suite.dotMeta, suite.settings, suite.template)
}

func (suite *DevelopTest) TestInstall() {
	suite.dotMeta.On("ReadMetaYml").Return(MetaYmlMock())
	suite.runner.On("Run", "docker", contains("build . --tag name")).Return()
	suite.develop.Install(false)
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestInstallNoCache() {
	suite.dotMeta.On("ReadMetaYml").Return(MetaYmlMock())
	suite.runner.On("Run", "docker", contains("build . --no-cache --tag name")).Return()
	suite.develop.Install(true)
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestEnter() {
	suite.dotMeta.On("ReadMetaYml").Return(MetaYmlMock())
	suite.runner.On("Run", "docker", contains("run -it name sh")).Return()
	suite.develop.Enter()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestStage() {
	metaYml := &meta.MetaYml{"name", "golang"}
	suite.dotMeta.On("ReadMetaYml").Return(metaYml)
	suite.log.On("Verbose", mock.Anything, mock.Anything)
	languageYml := map[string][]string{"build": []string{"cmd1", "cmd2 {{.Name}}"}}
	suite.settings.On("ReadLanguageYml", "golang").Return(&meta.LanguageYml{languageYml})
	suite.template.On("ExecuteOnString", "cmd1", mock.Anything).Return("cmd1")
	suite.template.On("ExecuteOnString", "cmd2 {{.Name}}", mock.Anything).Return("cmd2")
	suite.settings.On("Translation", metaYml).Return(&meta.Translation{*SettingsYmlMock(), *metaYml})

	suite.runner.On("Run", "docker", contains("run name cmd1")).Return()
	suite.runner.On("Run", "docker", contains("run name cmd2")).Return()
	suite.develop.Stage("build", true)
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestUpload() {
	suite.dotMeta.On("ReadMetaYml").Return(MetaYmlMock())
	suite.settings.On("ReadSettingsYml").Return(SettingsYmlMock())
	suite.runner.On("Run", "docker", contains("tag name url/namespace/name"))
	suite.runner.On("Run", "docker", contains("push url/namespace/name"))
	suite.develop.Upload()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestLogin() {
	suite.settings.On("ReadSettingsYml").Return(SettingsYmlMock())
	suite.runner.On("Run", "docker", contains("login -u user url")).Return()
	suite.develop.Login()
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
	suite.dotMeta.On("ReadMetaYml").Return(MetaYmlMock())
	suite.runner.On("Run", "docker", contains("run name cmd")).Return()
	suite.develop.Run([]string{"cmd"}, true)
	suite.runner.AssertExpectations(suite.T())
}

func TestDevelopSuite(t *testing.T) {
	suite.Run(t, new(DevelopTest))
}
