package meta

import (
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

//go:generate mockery -name=IUtil
type IUtil interface {
	ReadFile(path string) []byte
	ReadYml(path string, dataStruct interface{}) interface{}
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

type Util struct {
	log ILog
}

func NewUtil(log ILog) IUtil {
	return &Util{log}
}

func (self *Util) ReadFile(path string) []byte {
	content, err := ioutil.ReadFile(self.Expand(path))
	if err != nil {
		self.log.Fatal("Error in ReadFile:", err)
	}
	return content
}

func (self *Util) ReadYml(path string, dataStruct interface{}) interface{} {
	if err := yaml.Unmarshal(self.ReadFile(path), dataStruct); err != nil {
		self.log.Fatal("Error in ReadYml:", err)
	}
	return dataStruct
}

func (self *Util) ChangeDir(dir string) {
	if err := os.Chdir(self.Expand(dir)); err != nil {
		self.log.Fatal("Error in ChangeDir:", err)
	}
}

func (self *Util) GetCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		self.log.Fatal("Error in GetCwd:", err)
	}
	return cwd
}

func (self *Util) Exists(path string) bool {
	if _, err := os.Stat(self.Expand(path)); err == nil {
		return true
	}
	return false
}

func (self *Util) Mkdir(path string) {
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		self.log.Fatal("Error in Mkdir:", err)
	}
}

func (self *Util) Rename(from, to string) {
	if err := os.Rename(self.Expand(from), self.Expand(to)); err != nil {
		self.log.Fatal("Error in Rename:", err)
	}
}

func (self *Util) Walk(root string, walkFn filepath.WalkFunc) {
	if err := filepath.Walk(self.Expand(root), walkFn); err != nil {
		self.log.Fatal("Error in Walk:", err)
	}
}

func (self *Util) IsFile(path string) bool {
	fi, err := os.Stat(self.Expand(path))
	if err != nil {
		self.log.Fatal("Error in IsFile:", err)
	}
	return fi.Mode().IsRegular()
}

func (self *Util) CreateFile(filePath string) *os.File {
	os.MkdirAll(filepath.Dir(self.Expand(filePath)), os.ModePerm)
	f, err := os.Create(self.Expand(filePath))
	if err != nil {
		self.log.Fatal("Error in CreateFile:", err)
	}
	return f
}

func (self *Util) Expand(path string) string {
	expanded, err := homedir.Expand(path)
	if err != nil {
		self.log.Fatal("Error in expand:", err)
	}
	return expanded
}
