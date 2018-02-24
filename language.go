package main

import (
	"log"
)

type Language interface {
	Install()
	Build()
	Test()
	Lint()
	Coverage()
	CI()
	Run(args []string)
}

func NewLanguage(root *Root, config *Config) Language {
	switch config.Language {
	case "python":
		return NewPython(root, config)
	case "golang":
		return NewGolang(root, config)
	default:
		log.Fatal(config.Language, "not supported!")
		return NewPython(root, config)
	}
}
