package meta

import (
	"fmt"
	"reflect"
)

type ISetup interface {
	Run()
}

//go:generate mockery -name=ISetup
type Setup struct {
	util     IUtil
	settings ISettings
}

func NewSetup(util IUtil, settings ISettings) ISetup {
	return &Setup{util, settings}
}

func (self *Setup) Run() {
	settings := self.settings.ReadSettingsYml()
	settingsType := reflect.TypeOf(*settings)
	settingsValue := reflect.ValueOf(settings)

	for i := 0; i < settingsType.NumField(); i++ {
		current := settingsValue.Elem().Field(i).String()
		tag := settingsType.Field(i).Tag.Get("question")
		value := self.util.Input(fmt.Sprint("%s (%s): ", tag, current))
		if len(value) > 0 {
			settingsValue.Elem().Field(i).SetString(value)
		}
	}

	self.settings.WriteSettingsYml(settings)
}
