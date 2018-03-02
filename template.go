package main

import (
	"bytes"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type ITemplate interface {
	Create(metaYml *MetaYml)
}

type Template struct {
	walkRoot    string
	util        IUtil
	translation struct {
		MetaYml
		SettingsYml
	}
}

func NewTemplate(settings *Settings, util IUtil) ITemplate {
	t := new(Template)
	t.util = util
	t.translation.SettingsYml = *settings.Settings
	return t
}

func (t *Template) Create(metaYml *MetaYml) {
	t.translation.MetaYml = *metaYml
	t.walkRoot = path.Join(t.util.Expand("~"), ".meta", "templates", metaYml.Language)
	t.util.Walk(t.walkRoot, t.walkFn)
}

func (t *Template) walkFn(walkPath string, info os.FileInfo, err error) error {
	if t.util.IsFile(walkPath) {
		targetPathOutput := new(bytes.Buffer)
		relativeSoucePath, _ := filepath.Rel(t.walkRoot, walkPath)
		relativeTargetPath := path.Join(t.translation.MetaYml.Name, relativeSoucePath)
		template.Must(template.New("Filename").Parse(relativeTargetPath)).Execute(targetPathOutput, t.translation)
		newFile := t.util.CreateFile(targetPathOutput.String())
		template.Must(template.ParseFiles(walkPath)).Execute(newFile, t.translation)
		newFile.Close()
	}
	return nil
}
