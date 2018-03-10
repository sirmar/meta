package meta_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"meta/meta"
	"meta/meta/mocks"
	"os"
	"path"
	"testing"
)

type UtilTest struct {
	suite.Suite
	tmp  string
	log  *mocks.ILog
	util meta.IUtil
}

func (suite *UtilTest) SetupTest() {
	suite.tmp, _ = ioutil.TempDir("", "util_test")
	suite.log = new(mocks.ILog)
	suite.util = meta.NewUtil(suite.log)
}

func (suite *UtilTest) TearDownTest() {
	os.RemoveAll(suite.tmp)
}

func (suite *UtilTest) TestReadWriteFileAndExpand() {
	file := path.Join(suite.tmp, "file")
	suite.False(suite.util.Exists(file))
	suite.util.WriteFile(file, []byte("content"))

	suite.Equal([]byte("content"), suite.util.ReadFile(file))

	suite.log.On("Fatal", "Error in ReadFile:", mock.Anything)
	suite.util.ReadFile(suite.tmp)
	suite.log.AssertExpectations(suite.T())
}

func (suite *UtilTest) TestChangeDirAndGetCwd() {
	suite.NotEqual(suite.util.GetCwd(), suite.tmp)
	suite.util.ChangeDir(suite.tmp)
	suite.Equal(suite.util.GetCwd(), suite.tmp)
}

func (suite *UtilTest) TestRenameAndExists() {
	from := path.Join(suite.tmp, "from")
	to := path.Join(suite.tmp, "to")
	ioutil.WriteFile(from, []byte(""), os.ModePerm)

	suite.util.Rename(from, to)
	suite.True(suite.util.Exists(to))
	suite.False(suite.util.Exists(from))

	suite.log.On("Fatal", "Error in Rename:", mock.Anything)
	suite.util.Rename(from, to)
	suite.log.AssertExpectations(suite.T())
}

func (suite *UtilTest) TestMkdir() {
	dir := path.Join(suite.tmp, "newdir")
	suite.False(suite.util.Exists(dir))
	suite.util.Mkdir(dir)
	suite.True(suite.util.Exists(dir))

	suite.log.On("Fatal", "Error in Mkdir:", mock.Anything)
	wayout := path.Join(suite.tmp, "way", "out")
	suite.util.Mkdir(wayout)
	suite.log.AssertExpectations(suite.T())
}

func (suite *UtilTest) TestWalk() {
	walk := "not visited"
	suite.util.Walk(suite.tmp, func(_ string, _ os.FileInfo, _ error) error {
		walk = "visited"
		return nil
	})
	suite.Equal("visited", walk)
}

func (suite *UtilTest) TestIsFile() {
	file := path.Join(suite.tmp, "file")
	ioutil.WriteFile(file, []byte(""), os.ModePerm)

	suite.True(suite.util.IsFile(file))
	suite.False(suite.util.IsFile(suite.tmp))
}

func (suite *UtilTest) TestCreateFile() {
	path := path.Join(suite.tmp, "file_name")
	file := suite.util.CreateFile(path)
	suite.Contains(file.Name(), "file_name")
}

func TestUtilSuite(t *testing.T) {
	suite.Run(t, new(UtilTest))
}
