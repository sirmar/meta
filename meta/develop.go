package meta

import (
	"fmt"
	"strings"
)

//go:generate mockery -name=IDevelop
type IDevelop interface {
	Install()
	Enter()
	Stage(stage string, imageOnly bool)
	Upload()
	Run(args []string, imageOnly bool)
	Login()
}

type Develop struct {
	util     IUtil
	runner   IRunner
	dotMeta  IDotMeta
	settings ISettings
	template ITemplate
}

func NewDevelop(util IUtil, runner IRunner, dotMeta IDotMeta, settings ISettings, template ITemplate) IDevelop {
	return &Develop{util, runner, dotMeta, settings, template}
}

func (self *Develop) Install() {
	metaYml := self.dotMeta.ReadMetaYml()
	self.runner.Run("docker", []string{"build", ".", "--tag", metaYml.Name})
}

func (self *Develop) Enter() {
	metaYml := self.dotMeta.ReadMetaYml()
	self.runner.Run("docker", []string{"run", "-it", metaYml.Name, "sh"})
}

func (self *Develop) Stage(stage string, imageOnly bool) {
	metaYml := self.dotMeta.ReadMetaYml()
	for _, cmdTemplate := range self.settings.ReadLanguageYml(metaYml.Language).Stage(stage) {
		cmd := self.template.ExecuteOnString(cmdTemplate, self.settings.Translation(metaYml))
		parts := strings.Split(cmd, " ")
		args := append(self.baseArgs(metaYml, imageOnly), parts...)
		self.runner.Run("docker", args)
	}
}

func (self *Develop) Upload() {
	metaYml := self.dotMeta.ReadMetaYml()
	self.runner.Run("docker", []string{"push", metaYml.Name})
}

func (self *Develop) Run(args []string, imageOnly bool) {
	metaYml := self.dotMeta.ReadMetaYml()
	self.runner.Run("docker", append(self.baseArgs(metaYml, imageOnly), args...))
}

func (self *Develop) baseArgs(metaYml *MetaYml, imageOnly bool) []string {
	if imageOnly {
		return []string{"run", metaYml.Name}
	}
	return []string{"run", "-v", self.volume(metaYml), metaYml.Name}
}

func (self *Develop) volume(metaYml *MetaYml) string {
	root := self.util.GetCwd()
	if metaYml.Language == "golang" {
		return fmt.Sprintf("%s:/go/src/%s", root, metaYml.Name)
	}
	return fmt.Sprintf("%s:/usr/src/%s", root, metaYml.Name)
}

func (self *Develop) Login() {
	settingsYml := self.settings.ReadSettingsYml()
	self.runner.Run("docker", []string{"login", "-u", settingsYml.DockerUser, settingsYml.DockerRegistry})
}
