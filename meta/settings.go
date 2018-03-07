package meta

import (
	"fmt"
)

type CreateYml struct {
	Author string
	Email  string
}

type VerifyYml struct {
	RequiredFiles map[string][]string `yaml:"requiredFiles,flow"`
}

type LanguageYml struct {
	Stages map[string][]string `yaml:"stages,flow"`
}

type Translation struct {
	CreateYml
	MetaYml
}

func (self *LanguageYml) Stage(stage string) []string {
	value, _ := self.Stages[stage]
	return value
}

//go:generate mockery -name=ISettings
type ISettings interface {
	ReadCreateYml() *CreateYml
	ReadVerifyYml() *VerifyYml
	ReadLanguageYml(language string) *LanguageYml
	Translation(metaYml *MetaYml) interface{}
}

type Settings struct {
	util IUtil
}

func NewSettings(util IUtil) ISettings {
	return &Settings{util}
}

func (self *Settings) ReadCreateYml() *CreateYml {
	return self.util.ReadYml("~/.meta/create.yml", new(CreateYml)).(*CreateYml)
}

func (self *Settings) ReadVerifyYml() *VerifyYml {
	return self.util.ReadYml("~/.meta/verify.yml", new(VerifyYml)).(*VerifyYml)
}

func (self *Settings) ReadLanguageYml(language string) *LanguageYml {
	path := fmt.Sprintf("~/.meta/%s.yml", language)
	return self.util.ReadYml(path, new(LanguageYml)).(*LanguageYml)
}

func (self *Settings) Translation(metaYml *MetaYml) interface{} {
	return &Translation{*self.ReadCreateYml(), *metaYml}
}
