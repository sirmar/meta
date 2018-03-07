package meta

import (
	"path"
	"strings"
)

//go:generate mockery -name=IVerify
type IVerify interface {
	Files()
}

type Verify struct {
	util     IUtil
	settings ISettings
	log      ILog
}

func NewVerify(util IUtil, settings ISettings, log ILog) IVerify {
	return &Verify{util, settings, log}
}

func (self *Verify) Files() {
	for file, contentList := range self.settings.ReadVerifyYml().RequiredFiles {
		filePath := path.Join(self.util.GetCwd(), file)
		if !self.util.Exists(filePath) {
			self.log.Println("Missing file:", file)
			continue
		}
		content := self.util.ReadFile(filePath)
		for _, reqContent := range contentList {
			if !strings.Contains(string(content[:]), reqContent) {
				self.log.Println("File", file, "is missing content", reqContent)
			}
		}
	}

}
