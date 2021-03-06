// Code generated by mockery v1.0.0
package mocks

import meta "meta/meta"
import mock "github.com/stretchr/testify/mock"

// ISettings is an autogenerated mock type for the ISettings type
type ISettings struct {
	mock.Mock
}

// ReadLanguageYml provides a mock function with given fields: language
func (_m *ISettings) ReadLanguageYml(language string) *meta.LanguageYml {
	ret := _m.Called(language)

	var r0 *meta.LanguageYml
	if rf, ok := ret.Get(0).(func(string) *meta.LanguageYml); ok {
		r0 = rf(language)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*meta.LanguageYml)
		}
	}

	return r0
}

// ReadSettingsYml provides a mock function with given fields:
func (_m *ISettings) ReadSettingsYml() *meta.SettingsYml {
	ret := _m.Called()

	var r0 *meta.SettingsYml
	if rf, ok := ret.Get(0).(func() *meta.SettingsYml); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*meta.SettingsYml)
		}
	}

	return r0
}

// ReadVerifyYml provides a mock function with given fields:
func (_m *ISettings) ReadVerifyYml() *meta.VerifyYml {
	ret := _m.Called()

	var r0 *meta.VerifyYml
	if rf, ok := ret.Get(0).(func() *meta.VerifyYml); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*meta.VerifyYml)
		}
	}

	return r0
}

// Translation provides a mock function with given fields: metaYml
func (_m *ISettings) Translation(metaYml *meta.MetaYml) *meta.Translation {
	ret := _m.Called(metaYml)

	var r0 *meta.Translation
	if rf, ok := ret.Get(0).(func(*meta.MetaYml) *meta.Translation); ok {
		r0 = rf(metaYml)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*meta.Translation)
		}
	}

	return r0
}

// WriteSettingsYml provides a mock function with given fields: settingsYml
func (_m *ISettings) WriteSettingsYml(settingsYml *meta.SettingsYml) {
	_m.Called(settingsYml)
}
