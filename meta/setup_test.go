package meta_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"testing"
)

type SetupTest struct {
	suite.Suite
	util     *mocks.IUtil
	settings     *mocks.ISettings
	setup meta.ISetup
}

func (suite *SetupTest) SetupTest() {
	suite.util = new(mocks.IUtil)
	suite.settings = new(mocks.ISettings)
	suite.setup = meta.NewSetup(suite.util, suite.settings)
}

func (suite *SetupTest) TestSetValuesWritten() {
	suite.settings.On("ReadSettingsYml").Return(&meta.SettingsYml{"author", "email", "url", "namespace", "user"})
	suite.util.On("Input", mock.Anything).Return("input written")
	suite.settings.On("WriteSettingsYml", &meta.SettingsYml{
		"input written", "input written", "input written", "input written", "input written"})
	suite.setup.Run()
	suite.settings.AssertExpectations(suite.T())
}

func (suite *SetupTest) TestKeepPreviousValueWhenWritingNothing() {
	suite.settings.On("ReadSettingsYml").Return(&meta.SettingsYml{"author", "email", "url", "namespace", "user"})
	suite.util.On("Input", mock.Anything).Return("")
	suite.settings.On("WriteSettingsYml", &meta.SettingsYml{"author", "email", "url", "namespace", "user"})
	suite.setup.Run()
	suite.settings.AssertExpectations(suite.T())
}

func TestSetupSuite(t *testing.T) {
	suite.Run(t, new(SetupTest))
}
