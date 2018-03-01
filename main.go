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

	meta := NewMeta()
	settings := NewSettings()
	runner := NewDockerRunner(meta)
	language := NewLanguage(runner, meta)
	template := NewTemplate(settings)

	NewParser(language, template).Run()
}
