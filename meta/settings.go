package meta

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
	Settings *SettingsYml
	Python   *LanguageYml
	Golang   *LanguageYml
	util     IUtil
}

func NewSettings(util IUtil) *Settings {
	s := new(Settings)
	s.util = util
	return s
}

func (s *Settings) Read() *Settings {
	s.Settings = s.util.ReadYml("~/.meta/settings.yml", new(SettingsYml)).(*SettingsYml)
	s.Python = s.util.ReadYml("~/.meta/python.yml", new(LanguageYml)).(*LanguageYml)
	s.Golang = s.util.ReadYml("~/.meta/golang.yml", new(LanguageYml)).(*LanguageYml)
	return s
}
