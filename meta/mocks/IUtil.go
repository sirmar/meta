// Code generated by mockery v1.0.0
package mocks

import filepath "path/filepath"

import mock "github.com/stretchr/testify/mock"
import os "os"

// IUtil is an autogenerated mock type for the IUtil type
type IUtil struct {
	mock.Mock
}

// ChangeDir provides a mock function with given fields: dir
func (_m *IUtil) ChangeDir(dir string) {
	_m.Called(dir)
}

// CreateFile provides a mock function with given fields: path
func (_m *IUtil) CreateFile(path string) *os.File {
	ret := _m.Called(path)

	var r0 *os.File
	if rf, ok := ret.Get(0).(func(string) *os.File); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
		}
	}

	return r0
}

// Exists provides a mock function with given fields: path
func (_m *IUtil) Exists(path string) bool {
	ret := _m.Called(path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Expand provides a mock function with given fields: path
func (_m *IUtil) Expand(path string) string {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetCwd provides a mock function with given fields:
func (_m *IUtil) GetCwd() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Input provides a mock function with given fields: text
func (_m *IUtil) Input(text string) string {
	ret := _m.Called(text)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(text)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsFile provides a mock function with given fields: path
func (_m *IUtil) IsFile(path string) bool {
	ret := _m.Called(path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Mkdir provides a mock function with given fields: path
func (_m *IUtil) Mkdir(path string) {
	_m.Called(path)
}

// Mode provides a mock function with given fields: path
func (_m *IUtil) Mode(path string) os.FileMode {
	ret := _m.Called(path)

	var r0 os.FileMode
	if rf, ok := ret.Get(0).(func(string) os.FileMode); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(os.FileMode)
	}

	return r0
}

// ReadFile provides a mock function with given fields: path
func (_m *IUtil) ReadFile(path string) []byte {
	ret := _m.Called(path)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// ReadYml provides a mock function with given fields: path, dataStruct
func (_m *IUtil) ReadYml(path string, dataStruct interface{}) interface{} {
	ret := _m.Called(path, dataStruct)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, interface{}) interface{}); ok {
		r0 = rf(path, dataStruct)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Rename provides a mock function with given fields: from, to
func (_m *IUtil) Rename(from string, to string) {
	_m.Called(from, to)
}

// Stat provides a mock function with given fields: path
func (_m *IUtil) Stat(path string) os.FileInfo {
	ret := _m.Called(path)

	var r0 os.FileInfo
	if rf, ok := ret.Get(0).(func(string) os.FileInfo); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(os.FileInfo)
		}
	}

	return r0
}

// Walk provides a mock function with given fields: root, walkFn
func (_m *IUtil) Walk(root string, walkFn filepath.WalkFunc) {
	_m.Called(root, walkFn)
}

// WriteFile provides a mock function with given fields: path, content
func (_m *IUtil) WriteFile(path string, content []byte) {
	_m.Called(path, content)
}

// WriteYml provides a mock function with given fields: path, dataStruct
func (_m *IUtil) WriteYml(path string, dataStruct interface{}) {
	_m.Called(path, dataStruct)
}
