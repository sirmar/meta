package main

import "meta/meta"

func main() {
	log := meta.NewLog()
	util := meta.NewUtil(log)
	dotMeta := meta.NewDotMeta(util)
	settings := meta.NewSettings(util).Read()
	runner := meta.NewDockerRunner(log)
	language := meta.NewLanguage(runner, dotMeta)
	template := meta.NewTemplate(settings, util)
	meta.NewParser(language, template, log).Run()

}
