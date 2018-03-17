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
	suite.runner.On("Run", "docker", contains("build . --tag name"))
	suite.develop.Install(false)
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestInstallNoCache() {
	suite.dotMeta.On("ReadMetaYml").Return(MetaYmlMock())
	suite.runner.On("Run", "docker", contains("build . --no-cache --tag name"))
	suite.develop.Install(true)
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestEnter() {
	suite.dotMeta.On("ReadMetaYml").Return(MetaYmlMock())
	suite.runner.On("Run", "docker", contains("run -it name sh"))
	suite.develop.Enter()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestStage() {
	suite.dotMeta.On("ReadMetaYml").Return(MetaYmlMock())
	suite.log.On("Verbose", mock.Anything, mock.Anything)
	suite.settings.On("ReadLanguageYml", "language").Return(LanguageYmlMock())
	suite.template.On("ExecuteOnString", "build cmd", mock.Anything).Return("build cmd")
	suite.settings.On("Translation", mock.Anything).Return(TranslationMock())

	suite.runner.On("Run", "docker", contains("run name build cmd"))
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
	suite.runner.On("Run", "docker", contains("login -u user url"))
	suite.develop.Login()
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestRun() {
	suite.dotMeta.On("ReadMetaYml").Return(MetaYmlMock())
	suite.settings.On("ReadLanguageYml", "language").Return(LanguageYmlMock())
	suite.settings.On("Translation", mock.Anything).Return(TranslationMock())
	suite.template.On("ExecuteOnString", "/root:/usr/src/name", mock.Anything).Return("/root:/usr/src/name")
	suite.runner.On("Run", "docker", contains("run -v /root:/usr/src/name name cmd"))
	suite.develop.Run([]string{"cmd"}, false)
	suite.runner.AssertExpectations(suite.T())
}

func (suite *DevelopTest) TestRunWithImageOnly() {
	suite.dotMeta.On("ReadMetaYml").Return(MetaYmlMock())
	suite.runner.On("Run", "docker", contains("run name cmd"))
	suite.develop.Run([]string{"cmd"}, true)
	suite.runner.AssertExpectations(suite.T())
}

func TestDevelopSuite(t *testing.T) {
	suite.Run(t, new(DevelopTest))
}
