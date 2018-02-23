package main

import (
	"os"
	"log"
	"path"
	"path/filepath"
)

type Root struct {
	Root string
	Cwd string
}

func NewRoot() *Root {
	currentDir, _ := os.Getwd()
	return &Root{find(currentDir), currentDir}
}

func (r *Root) MoveToRoot() {
	os.Chdir(r.Root)
}

func (r *Root) MoveToCwd() {
	os.Chdir(r.Cwd)
}

func find(currentDir string) string {
	metaFilePath := path.Join(currentDir, ".meta", "meta.yml")
	if _, err := os.Stat(metaFilePath); err == nil {
		return currentDir;
	}
	if currentDir == "/" {
		log.Fatal("Meta not found!")
	}
	return find(filepath.Dir(currentDir))
}
