package meta

import (
	"os"
	"path"
	"path/filepath"
)

//go:generate mockery -name=ICreate
type ICreate interface {
	Template(language, name string)
}

type Create struct {
	util     IUtil
	settings ISettings
	template ITemplate
}

func NewCreate(util IUtil, settings ISettings, template ITemplate) ICreate {
	return &Create{util, settings, template}
}

func (self *Create) Template(language, name string) {
	walkRoot := path.Join(self.util.Expand("~"), ".meta", "templates", language)
	translation := self.settings.Translation(&MetaYml{language, name})
	self.util.Walk(walkRoot, func(walkPath string, _ os.FileInfo, _ error) error {
		if self.util.IsFile(walkPath) {
			soucePathTemplate, _ := filepath.Rel(walkRoot, walkPath)
			targetPathTemplate := path.Join(name, soucePathTemplate)
			targetPath := self.template.ExecuteOnString(targetPathTemplate, translation)
			self.template.ExecuteOnFile(walkPath, targetPath, translation)
		}
		return nil
	})
}
