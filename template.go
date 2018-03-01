package main

type ITemplate interface {
	Create(name string, language string)
}

type Template struct {
	Settings *Settings
}

func NewTemplate(settings *Settings) ITemplate {
	return &Template{settings}
}

func (t *Template) Create(name string, language string) {
}
