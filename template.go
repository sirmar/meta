package main

type Template struct {
	language string
	name     string
}

func NewTempate(language string, name string) *Template {
	return &Template{language, name}
}

func (t *Template) Create() {
}
