package meta_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"testing"
)

type VerifyTest struct {
	suite.Suite
	util     *mocks.IUtil
	settings *mocks.ISettings
	log      *mocks.ILog
	verify   meta.IVerify
}

func (suite *VerifyTest) SetupTest() {
	suite.util = new(mocks.IUtil)
	suite.settings = new(mocks.ISettings)
	suite.log = new(mocks.ILog)
	suite.verify = meta.NewVerify(suite.util, suite.settings, suite.log)
}

func (suite *VerifyTest) TestNoFilesToVerify() {
	suite.givenNoRequirements()

	suite.verify.Files()
	suite.log.AssertExpectations(suite.T()) // No logs
}

func (suite *VerifyTest) TestFileExists() {
	suite.givenRequirementOnFileWithContent("file", "")
	suite.givenFileWithContent("some content")

	suite.verify.Files()
	suite.log.AssertExpectations(suite.T()) // No logs
}

func (suite *VerifyTest) TestFileExistsWithContent() {
	suite.givenRequirementOnFileWithContent("file", "required content")
	suite.givenFileWithContent("some required content")

	suite.verify.Files()
	suite.log.AssertExpectations(suite.T()) // No logs
}

func (suite *VerifyTest) TestFileMissing() {
	suite.givenRequirementOnFileWithContent("file", "")
	suite.givenFileMissing("file")

	suite.log.On("Println", "Missing file:", "file").Return()
	suite.verify.Files()
	suite.log.AssertExpectations(suite.T())
}

func (suite *VerifyTest) TestFileExistsButMissRequiredContent() {
	suite.givenRequirementOnFileWithContent("file", "required content")
	suite.givenFileWithContent("missing content")

	suite.log.On("Println", "File", "file", "is missing content", "required content").Return()
	suite.verify.Files()
	suite.log.AssertExpectations(suite.T()) // No logs
}

func (suite *VerifyTest) givenNoRequirements() {
	suite.settings.On("ReadVerifyYml").Return(&meta.VerifyYml{map[string][]string{}})
}

func (suite *VerifyTest) givenRequirementOnFileWithContent(file, content string) {
	requiredFiles := map[string][]string{file: []string{content}}
	suite.settings.On("ReadVerifyYml").Return(&meta.VerifyYml{requiredFiles})
}

func (suite *VerifyTest) givenFileWithContent(content string) {
	suite.util.On("GetCwd").Return("/dir")
	suite.util.On("Exists", "/dir/file").Return(true)
	suite.util.On("ReadFile", "/dir/file").Return([]byte(content))
}

func (suite *VerifyTest) givenFileMissing(file string) {
	path := fmt.Sprintf("/dir/%s", file)
	suite.util.On("GetCwd").Return("/dir")
	suite.util.On("Exists", path).Return(false)
}

func TestVerifySuite(t *testing.T) {
	suite.Run(t, new(VerifyTest))
}
