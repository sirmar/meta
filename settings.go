package main

type SettingsYml struct {
	Author string
	Email  string
}

type LanguageYml struct {
	Build    struct{ cmds []string }
	Test     struct{ cmds []string }
	Lint     struct{ cmds []string }
	Coverage struct{ cmds []string }
	CI       struct{ cmds []string }
}

type Settings struct {
	Settings *SettingsYml
	Python   *LanguageYml
	Golang   *LanguageYml
}

func NewSettings() *Settings {
	return &Settings{
		ReadYml("~/.meta/general.yml", new(SettingsYml)).(*SettingsYml),
		ReadYml("~/.meta/python.yml", new(LanguageYml)).(*LanguageYml),
		ReadYml("~/.meta/golang.yml", new(LanguageYml)).(*LanguageYml),
	}
}
