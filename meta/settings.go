package meta

import (
	"fmt"
)

type SettingsYml struct {
	Author          string `yaml:"author" question:"Your name"`
	Email           string `yaml:"email" question:"Your email"`
	DockerRegistry  string `yaml:"docker-registry" question:"Url to docker registry"`
	DockerNamespace string `yaml:"docker-namespace" question:"Namespace used in docker registry"`
	DockerUser      string `yaml:"docker-user" question:"User name used when logging in to docker registry"`
}

type VerifyYml struct {
	RequiredFiles map[string][]string `yaml:"requiredFiles,flow"`
}

type LanguageYml struct {
	Stages map[string][]string `yaml:"stages,flow"`
}

type Translation struct {
	SettingsYml
	MetaYml
}

func (self *LanguageYml) Stage(stage string) []string {
	value, _ := self.Stages[stage]
	return value
}

//go:generate mockery -name=ISettings
type ISettings interface {
	WriteSettingsYml(settingsYml *SettingsYml)
	ReadSettingsYml() *SettingsYml
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

func (self *Settings) ReadSettingsYml() *SettingsYml {
	return self.util.ReadYml("~/.meta/settings.yml", new(SettingsYml)).(*SettingsYml)
}

func (self *Settings) ReadVerifyYml() *VerifyYml {
	return self.util.ReadYml("~/.meta/verify.yml", new(VerifyYml)).(*VerifyYml)
}

func (self *Settings) ReadLanguageYml(language string) *LanguageYml {
	path := fmt.Sprintf("~/.meta/%s.yml", language)
	return self.util.ReadYml(path, new(LanguageYml)).(*LanguageYml)
}

func (self *Settings) Translation(metaYml *MetaYml) interface{} {
	return &Translation{*self.ReadSettingsYml(), *metaYml}
}

func (self *Settings) WriteSettingsYml(settingsYml *SettingsYml) {
	self.util.WriteYml("~/.meta/settings.yml", settingsYml)
}
