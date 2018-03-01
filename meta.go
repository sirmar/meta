package main

import (
	"os"
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
}

func (m *Meta) MoveToRoot() {
	os.Chdir(m.Root)
}

func NewMeta() *Meta {
	m := new(Meta)
	currentDir, _ := os.Getwd()
	m.Root = find(currentDir)
	m.Cwd = currentDir

	if m.Found() {
		m.Meta = ReadYml(".meta/meta.yml", new(MetaYml)).(*MetaYml)
	}

	return m
}

func (m *Meta) MoveToCwd() {
	os.Chdir(m.Cwd)
}

func (m *Meta) Found() bool {
	return m.Root != "Not found"
}

func find(currentDir string) string {
	metaFilePath := path.Join(currentDir, ".meta", "meta.yml")
	if _, err := os.Stat(metaFilePath); err == nil {
		return currentDir
	}
	if currentDir == "/" {
		return "Not found"
	}
	return find(filepath.Dir(currentDir))
}
