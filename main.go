package main

import (
	"fmt"
	"log"
)

type logWriter struct {
}

func (w logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print("Meta: ", string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	util := NewUtil()
	meta := NewMeta(util)
	settings := NewSettings(util).read()
	runner := NewDockerRunner(meta)
	language := NewLanguage(runner, meta)
	template := NewTemplate(settings, util)
	NewParser(language, template).Run()

}
