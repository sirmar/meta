package meta

import (
	"fmt"
)

type SettingsYml struct {
	Author string
	Email  string
}

type LanguageYml struct {
	Build    []string `yaml:"build,flow"`
	Test     []string `yaml:"test,flow"`
	Lint     []string `yaml:"lint,flow"`
	Coverage []string `yaml:"coverage,flow"`
	CI       []string `yaml:"ci,flow"`
}

type Settings struct {
	util IUtil
}

func NewSettings(util IUtil) *Settings {
	return &Settings{util}
}

func (s *Settings) ReadSettings() *SettingsYml {
	return s.util.ReadYml("~/.meta/settings.yml", new(SettingsYml)).(*SettingsYml)
}

func (s *Settings) ReadLanguage(language string) *LanguageYml {
	path := fmt.Sprintf("~/.meta/%s.yml", language)
	return s.util.ReadYml(path, new(LanguageYml)).(*LanguageYml)
}
