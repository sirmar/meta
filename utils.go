package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func ReadYml(path string, dataStruct interface{}) interface{} {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(content, dataStruct); err != nil {
		log.Fatal(err)
	}
	return dataStruct
}
