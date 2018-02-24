package main

import (
	"fmt"
	"log"
)

type logWriter struct {
}

func (w logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	root := NewRoot()
	config := NewConfig()
	language := NewLanguage(root, config)
	NewParser(language)
}
