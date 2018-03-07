package meta

import (
	"path"
)

type MetaYml struct {
	Name     string
	Language string
}

//go:generate mockery -name=IDotMeta
type IDotMeta interface {
	ReadMetaYml() *MetaYml
	MoveToRoot()
}

type DotMeta struct {
	util IUtil
}

func NewDotMeta(util IUtil) IDotMeta {
	return &DotMeta{util}
}

func (self *DotMeta) ReadMetaYml() *MetaYml {
	return self.util.ReadYml(path.Join(".meta", "meta.yml"), new(MetaYml)).(*MetaYml)
}

func (self *DotMeta) MoveToRoot() {
	if root := self.find(self.util.GetCwd()); root != "" {
		self.util.ChangeDir(root)
	}
}

func (self *DotMeta) find(currentDir string) string {
	metaFilePath := path.Join(currentDir, ".meta", "meta.yml")
	if self.util.Exists(metaFilePath) {
		return currentDir
	} else if currentDir == "/" {
		return ""
	}
	return self.find(path.Dir(currentDir))
}
