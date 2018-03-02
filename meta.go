package main

import (
	"path"
	"path/filepath"
)

type MetaYml struct {
	Name     string
	Language string
}

type Meta struct {
	Meta *MetaYml
	Root string
	Cwd  string
	util IUtil
}

func (m *Meta) MoveToRoot() {
	m.util.ChangeDir(m.Root)
}

func NewMeta(util IUtil) *Meta {
	m := new(Meta)
	m.util = util
	currentDir := util.GetCwd()
	m.Root = m.find(currentDir)
	m.Cwd = currentDir

	if m.Found() {
		m.Meta = util.ReadYml(".meta/meta.yml", new(MetaYml)).(*MetaYml)
	}

	return m
}

func (m *Meta) MoveToCwd() {
	m.util.ChangeDir(m.Cwd)
}

func (m *Meta) Found() bool {
	return m.Root != "Not found"
}

func (m *Meta) find(currentDir string) string {
	metaFilePath := path.Join(currentDir, ".meta", "meta.yml")
	if m.util.Exists(metaFilePath) {
		return currentDir
	} else if currentDir == "/" {
		return "Not found"
	}
	return m.find(filepath.Dir(currentDir))
}
