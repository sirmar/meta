package meta

import (
	"fmt"
	"strings"
)

type ICommand interface {
	Install()
	Enter()
	Run(args []string, imageOnly bool)
	Create(name, language string)
	Language(cmds []string, imageOnly bool)
}

type Command struct {
	runner   IRunner
	dotMeta  *DotMeta
	template ITemplate
}

func NewCommand(runner IRunner, dotMeta *DotMeta, template ITemplate) ICommand {
	return &Command{runner, dotMeta, template}
}

func (self *Command) Install() {
	self.runner.Run("docker", []string{"build", ".", "--tag", self.dotMeta.MetaYml.Name})
}

func (self *Command) Enter() {
	self.runner.Run("docker", []string{"run", "-it", self.dotMeta.MetaYml.Name, "sh"})
}

func (self *Command) Run(args []string, imageOnly bool) {
	self.runner.Run("docker", append(self.baseArgs(imageOnly), args...))
}

func (self *Command) Create(name, language string) {
	self.template.Create(name, language)
}

func (self *Command) Language(cmds []string, imageOnly bool) {
	for _, cmd := range cmds {
		parts := strings.Split(cmd, " ")
		args := append(self.baseArgs(imageOnly), parts...)
		self.runner.Run("docker", args)
	}
}

func (self *Command) baseArgs(imageOnly bool) []string {
	if imageOnly {
		return []string{"run", self.dotMeta.MetaYml.Name}
	}
	return []string{"run", "-v", self.volume(), self.dotMeta.MetaYml.Name}
}

func (self *Command) volume() string {
	if self.dotMeta.MetaYml.Language == "golang" {
		return fmt.Sprintf("%s:/go/src/%s", self.dotMeta.Root, self.dotMeta.MetaYml.Name)
	}
	return fmt.Sprintf("%s:/usr/src/%s", self.dotMeta.Root, self.dotMeta.MetaYml.Name)
}
