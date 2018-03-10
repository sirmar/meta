package meta_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"testing"
)

type SettingsTest struct {
	suite.Suite
	util     *mocks.IUtil
	settings meta.ISettings
}

func (suite *SettingsTest) SetupTest() {
	suite.util = new(mocks.IUtil)
	suite.settings = meta.NewSettings(suite.util)
}

func (suite *SettingsTest) TestReadSettingsYml() {
	suite.util.On("ReadYml", "~/.meta/settings.yml", mock.Anything).Return(
		&meta.SettingsYml{"author", "email", "url", "namespace", "user"})
	createYml := suite.settings.ReadSettingsYml()
	suite.Equal("author", createYml.Author)
	suite.Equal("email", createYml.Email)
	suite.Equal("url", createYml.DockerRegistry)
	suite.Equal("namespace", createYml.DockerNamespace)
	suite.Equal("user", createYml.DockerUser)
}

func (suite *SettingsTest) TestWriteSettingsYml() {
	suite.util.On("WriteYml", "~/.meta/settings.yml", mock.Anything).Return()
	suite.settings.WriteSettingsYml(&meta.SettingsYml{"author", "email", "url", "namespace", "user"})
	suite.util.AssertExpectations(suite.T())
}

func (suite *SettingsTest) TestReadVerifyYml() {
	content := map[string][]string{"file": []string{"content"}}
	suite.util.On("ReadYml", "~/.meta/verify.yml", mock.Anything).Return(&meta.VerifyYml{content})
	verifyYml := suite.settings.ReadVerifyYml()
	suite.Equal("content", verifyYml.RequiredFiles["file"][0])
}

func (suite *SettingsTest) TestReadLanguageYml() {
	content := map[string][]string{"build": []string{"build cmd"}}
	suite.util.On("ReadYml", "~/.meta/golang.yml", mock.Anything).Return(&meta.LanguageYml{content})
	languageYml := suite.settings.ReadLanguageYml("golang")
	suite.Equal([]string{"build cmd"}, languageYml.Stage("build"))
}

func (suite *SettingsTest) TestTranslation() {
	suite.util.On("ReadYml", "~/.meta/settings.yml", mock.Anything).Return(
		&meta.SettingsYml{"author", "email", "url", "namespace", "user"})
	translation := suite.settings.Translation(&meta.MetaYml{"name", "language"}).(*meta.Translation)
	suite.Equal("author", translation.Author)
	suite.Equal("email", translation.Email)
	suite.Equal("url", translation.DockerRegistry)
	suite.Equal("namespace", translation.DockerNamespace)
	suite.Equal("user", translation.DockerUser)
	suite.Equal("name", translation.Name)
	suite.Equal("language", translation.Language)
}

func TestSettingsSuite(t *testing.T) {
	suite.Run(t, new(SettingsTest))
}
