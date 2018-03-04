package meta

import (
	"bytes"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type ITemplate interface {
	Create(name, language string)
}

type Template struct {
	util        IUtil
	settingsYml *SettingsYml
}

func NewTemplate(util IUtil, settingsYml *SettingsYml) ITemplate {
	return &Template{util, settingsYml}
}

func (t *Template) Create(name, language string) {
	walkRoot := path.Join(t.util.Expand("~"), ".meta", "templates", language)
	translation := t.translation(name, language)

	t.util.Walk(walkRoot, func(walkPath string, info os.FileInfo, err error) error {
		if t.util.IsFile(walkPath) {
			targetPathOutput := new(bytes.Buffer)
			relativeSoucePath, _ := filepath.Rel(walkRoot, walkPath)
			relativeTargetPath := path.Join(name, relativeSoucePath)
			template.Must(template.New("Filename").Parse(relativeTargetPath)).Execute(targetPathOutput, translation)
			newFile := t.util.CreateFile(targetPathOutput.String())
			template.Must(template.ParseFiles(walkPath)).Execute(newFile, translation)
			newFile.Close()
		}
		return nil
	})
}

func (t *Template) translation(name, language string) interface{} {
	return struct {
		SettingsYml
		MetaYml
	}{
		*t.settingsYml,
		MetaYml{name, language},
	}
}
