package meta_test

import (
	"github.com/stretchr/testify/suite"
	"meta/meta"
	"meta/meta/mocks"
	"testing"
)

type TemplateTest struct {
	suite.Suite
	util     *mocks.IUtil
	template meta.ITemplate
}

func (suite *TemplateTest) SetupTest() {
	suite.util = new(mocks.IUtil)
	suite.template = meta.NewTemplate(suite.util)
}

func (suite *TemplateTest) TestExecuteOnString() {
	type Translation struct {
		Name string
	}

	suite.Equal("name", suite.template.ExecuteOnString("{{.Name}}", &Translation{"name"}))
	suite.Empty(suite.template.ExecuteOnString("{{.Missing}}", &Translation{"name"}))
}

func TestTemplateSuite(t *testing.T) {
	suite.Run(t, new(TemplateTest))
}
