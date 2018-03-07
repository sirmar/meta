package meta_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"path"
	"path/filepath"
	"testing"
)

type CreateTest struct {
	suite.Suite
	util     *mocks.IUtil
	settings *mocks.ISettings
	template *mocks.ITemplate
	create   meta.ICreate
}

func (suite *CreateTest) SetupTest() {
	suite.util = new(mocks.IUtil)
	suite.settings = new(mocks.ISettings)
	suite.template = new(mocks.ITemplate)
	suite.create = meta.NewCreate(suite.util, suite.settings, suite.template)
}

func (suite *CreateTest) TestDirectory() {
	suite.givenTranslation(&meta.MetaYml{"golang", "newpkg"})
	suite.givenDirectory()
	suite.create.Template("golang", "newpkg")
	suite.util.AssertExpectations(suite.T())
	suite.template.AssertExpectations(suite.T()) // No calls
}

func (suite *CreateTest) TestFilePathAndContentWillBeTemplateExpanded() {
	suite.givenTranslation(&meta.MetaYml{"golang", "newpkg"})
	suite.givenFile("file_{{.Name}}")
	suite.create.Template("golang", "newpkg")
	suite.util.AssertExpectations(suite.T())
	suite.template.AssertExpectations(suite.T())
}

func (suite *CreateTest) givenTranslation(metaYml *meta.MetaYml) {
	suite.settings.On("Translation", metaYml).Return(
		&meta.Translation{meta.CreateYml{"John", "john@email.com"}, *metaYml})
}

func (suite *CreateTest) givenDirectory() {
	suite.givenFileOrDirWithContent("dir", false)
}

func (suite *CreateTest) givenFile(file string) {
	suite.givenFileOrDirWithContent(file, true)
}

func (suite *CreateTest) givenFileOrDirWithContent(fileOrDir string, isFile bool) {
	suite.util.On("Expand", mock.Anything).Return("/home")
	suite.util.On("Walk", mock.Anything, mock.AnythingOfType("filepath.WalkFunc")).Return(nil).Run(func(args mock.Arguments) {
		walkRoot := args.String(0)
		walkFn := args.Get(1).(filepath.WalkFunc)
		walkPath := path.Join(walkRoot, fileOrDir)
		suite.util.On("IsFile", walkPath).Return(isFile)
		if isFile {
			suite.template.On("ExecuteOnString", "newpkg/file_{{.Name}}", mock.Anything).Return("newpkg/file_newpkg")
			suite.template.On("ExecuteOnFile", mock.Anything, "newpkg/file_newpkg", mock.Anything)
		}
		walkFn(walkPath, nil, nil)
	})
}

func TestCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateTest))
}
