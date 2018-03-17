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
	log  ILog
}

func NewTemplate(util IUtil, log ILog) ITemplate {
	return &Template{util, log}
}

func (self *Template) ExecuteOnString(input string, translation interface{}) string {
	output := new(bytes.Buffer)
	template.Must(template.New("String").Parse(input)).Execute(output, translation)
	self.log.Verbose("Expanding string", input, "to", output.String())
	return output.String()
}

func (self *Template) ExecuteOnFile(sourcePath, targetPath string, translation interface{}) {
	self.log.Verbose("Expanding template", sourcePath, "to", targetPath)
	newFile := self.util.CreateFile(targetPath)
	newFile.Chmod(self.util.Mode(sourcePath))
	template.Must(template.ParseFiles(sourcePath)).Execute(newFile, translation)
	newFile.Close()
}
