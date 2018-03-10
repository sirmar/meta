package meta

import (
	"bytes"
	"text/template"
)

//go:generate mockery -name=ITemplate
type ITemplate interface {
	ExecuteOnString(input string, translation interface{}) string
	ExecuteOnFile(sourcePath, targetPath string, translation interface{})
}

type Template struct {
	util IUtil
}

func NewTemplate(util IUtil) ITemplate {
	return &Template{util}
}

func (self *Template) ExecuteOnString(input string, translation interface{}) string {
	output := new(bytes.Buffer)
	template.Must(template.New("String").Parse(input)).Execute(output, translation)
	return output.String()
}

func (self *Template) ExecuteOnFile(sourcePath, targetPath string, translation interface{}) {
	newFile := self.util.CreateFile(targetPath)
	newFile.Chmod(self.util.Mode(sourcePath))
	template.Must(template.ParseFiles(sourcePath)).Execute(newFile, translation)
	newFile.Close()
}
