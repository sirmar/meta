package main

import (
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

type IUtil interface {
	ReadFile(path string) []byte
	ReadYml(path string, dataStruct interface{}) interface{}
	Cp(from, to string)
	ChangeDir(dir string)
	GetCwd() string
	Exists(path string) bool
	Mkdir(path string)
	Rename(from, to string)
	Walk(root string, walkFn filepath.WalkFunc)
	IsFile(path string) bool
	CreateFile(path string) *os.File
	Expand(path string) string
}

type Util struct{}

func NewUtil() IUtil {
	return &Util{}
}

func (u *Util) ReadFile(path string) []byte {
	content, err := ioutil.ReadFile(u.Expand(path))
	if err != nil {
		log.Fatal("Error in ReadFile:", err)
	}
	return content
}

func (u *Util) ReadYml(path string, dataStruct interface{}) interface{} {
	if err := yaml.Unmarshal(u.ReadFile(path), dataStruct); err != nil {
		log.Fatal("Error in ReadYml:", err)
	}
	return dataStruct
}

func (u *Util) Cp(from, to string) {
	if err := exec.Command("cp", "-r", u.Expand(from), u.Expand(to)).Run(); err != nil {
		log.Fatal("Error in CopyFile:", err)
	}
}

func (u *Util) ChangeDir(dir string) {
	if err := os.Chdir(u.Expand(dir)); err != nil {
		log.Fatal("Error in ChangeDir:", err)
	}
}

func (u *Util) GetCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error in GetCwd:", err)
	}
	return cwd
}

func (u *Util) Exists(path string) bool {
	if _, err := os.Stat(u.Expand(path)); err == nil {
		return true
	}
	return false
}

func (u *Util) Mkdir(path string) {
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		log.Fatal("Error in Mkdir:", err)
	}
}

func (u *Util) Rename(from, to string) {
	if err := os.Rename(u.Expand(from), u.Expand(to)); err != nil {
		log.Fatal("Error in Rename:", err)
	}
}

func (u *Util) Walk(root string, walkFn filepath.WalkFunc) {
	if err := filepath.Walk(u.Expand(root), walkFn); err != nil {
		log.Fatal("Error in Walk:", err)
	}
}

func (u *Util) IsFile(path string) bool {
	fi, err := os.Stat(u.Expand(path))
	if err != nil {
		log.Fatal("Error in IsFile:", err)
	}
	return fi.Mode().IsRegular()
}

func (u *Util) CreateFile(filePath string) *os.File {
	os.MkdirAll(path.Dir(u.Expand(filePath)), os.ModePerm)
	f, err := os.Create(u.Expand(filePath))
	if err != nil {
		log.Fatal("Error in CreateFile:", err)
	}
	return f
}

func (u *Util) Expand(path string) string {
	expanded, err := homedir.Expand(path)
	if err != nil {
		log.Fatal("Error in expand:", err)
	}
	return expanded
}
