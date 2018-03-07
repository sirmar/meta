package main

import "meta/meta"

func main() {
	log := meta.NewLog()
	util := meta.NewUtil(log)
	dotMeta := meta.NewDotMeta(util)
	settings := meta.NewSettings(util)
	template := meta.NewTemplate(util)
	runner := meta.NewRunner(log)
	create := meta.NewCreate(util, settings, template)
	verify := meta.NewVerify(util, settings, log)
	develop := meta.NewDevelop(util, runner, dotMeta, settings)
	command := meta.NewCommand(develop, create, verify)
	meta.NewParser(dotMeta, command, log).Run()
}