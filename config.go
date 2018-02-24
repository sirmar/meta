package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Name     string
	Language string
}

func NewConfig() *Config {
	data, err := ioutil.ReadFile(".meta/meta.yml")
	if err != nil {
		log.Fatal(err)
	}

	config := new(Config)
	if err := yaml.Unmarshal(data, config); err != nil {
		log.Fatal(err)
	}
	return config
}
