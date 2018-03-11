package meta_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"testing"
)

type DotMetaTest struct {
	suite.Suite
	util    *mocks.IUtil
	dotMeta meta.IDotMeta
}

func MetaYmlMock() *meta.MetaYml {
	return &meta.MetaYml{"name", "language"}
}

func (suite *DotMetaTest) SetupTest() {
	suite.util = new(mocks.IUtil)
	suite.dotMeta = meta.NewDotMeta(suite.util)
}

func (suite *DotMetaTest) TestReadMetaYml() {
	suite.util.On("ReadYml", ".meta/meta.yml", mock.Anything).Return(MetaYmlMock())
	metaYml := suite.dotMeta.ReadMetaYml()
	suite.Equal("name", metaYml.Name)
	suite.Equal("language", metaYml.Language)
}

func (suite *DotMetaTest) TestMoveToRootFindsRoot() {
	suite.util.On("GetCwd").Return("/current/dir")
	suite.util.On("Exists", "/current/dir/.meta/meta.yml").Return(false)
	suite.util.On("Exists", "/current/.meta/meta.yml").Return(true)

	suite.util.On("ChangeDir", "/current").Return()
	suite.dotMeta.MoveToRoot()
	suite.util.AssertExpectations(suite.T())
}

func (suite *DotMetaTest) TestMoveToRootNeverFindsRoot() {
	suite.util.On("GetCwd").Return("/dir")
	suite.util.On("Exists", mock.Anything).Return(false)

	suite.dotMeta.MoveToRoot()
	suite.util.AssertExpectations(suite.T())
}

func TestDotMetaSuite(t *testing.T) {
	suite.Run(t, new(DotMetaTest))
}
