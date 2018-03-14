package main

import "meta/meta"

func main() {
	log := meta.NewLog()
	util := meta.NewUtil(log)
	dotMeta := meta.NewDotMeta(util)
	settings := meta.NewSettings(util)
	template := meta.NewTemplate(util, log)
	runner := meta.NewRunner(log)
	setup := meta.NewSetup(util, settings)
	create := meta.NewCreate(util, settings, template)
	verify := meta.NewVerify(util, settings, log)
	release := meta.NewRelease(util, runner, log)
	develop := meta.NewDevelop(util, log, runner, dotMeta, settings, template)

	meta.NewParser(dotMeta, log, develop, create, verify, setup, release).Run()
}
