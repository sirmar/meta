package meta_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/mock"
	"meta/meta"
	"meta/meta/mocks"
	"testing"
)

type TemplateTest struct {
	suite.Suite
	util     *mocks.IUtil
	log     *mocks.ILog
	template meta.ITemplate
}

func (suite *TemplateTest) SetupTest() {
	suite.util = new(mocks.IUtil)
	suite.log = new(mocks.ILog)
	suite.template = meta.NewTemplate(suite.util, suite.log)
}

func (suite *TemplateTest) TestExecuteOnString() {
	type Translation struct {
		Name string
	}

	suite.log.On("Verbose", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	suite.Equal("name", suite.template.ExecuteOnString("{{.Name}}", &Translation{"name"}))
	suite.Empty(suite.template.ExecuteOnString("{{.Missing}}", &Translation{"name"}))
}

func TestTemplateSuite(t *testing.T) {
	suite.Run(t, new(TemplateTest))
}
