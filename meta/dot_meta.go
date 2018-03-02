package meta

import (
	"path"
	"path/filepath"
)

type MetaYml struct {
	Name     string
	Language string
}

type DotMeta struct {
	Meta *MetaYml
	Root string
	Util IUtil
}

func NewDotMeta(util IUtil) *DotMeta {
	m := new(DotMeta)
	m.Util = util
	currentDir := util.GetCwd()
	m.Root = m.find(currentDir)

	if m.Found() {
		m.Util.ChangeDir(m.Root)
		metaPath := path.Join(m.Root, ".meta", "meta.yml")
		m.Meta = util.ReadYml(metaPath, new(MetaYml)).(*MetaYml)
	}

	return m
}

func (m *DotMeta) Found() bool {
	return m.Root != "Not found"
}

func (m *DotMeta) find(currentDir string) string {
	metaFilePath := path.Join(currentDir, ".meta", "meta.yml")
	if m.Util.Exists(metaFilePath) {
		return currentDir
	} else if currentDir == "/" {
		return "Not found"
	}
	return m.find(filepath.Dir(currentDir))
}
