package main

import "meta/meta"

func main() {
	log := meta.NewLog()
	util := meta.NewUtil(log)
	dotMeta := meta.NewDotMeta(util)
	settings := meta.NewSettings(util)

	languageYml := settings.ReadLanguage(dotMeta.MetaYml.Language)
	settingsYml := settings.ReadSettings()

	runner := meta.NewRunner(log)
	template := meta.NewTemplate(util, settingsYml)
	command := meta.NewCommand(runner, dotMeta, template)
	meta.NewParser(languageYml, command, log).Run()
}
