// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// ICreate is an autogenerated mock type for the ICreate type
type ICreate struct {
	mock.Mock
}

// Template provides a mock function with given fields: language, name
func (_m *ICreate) Template(language string, name string) {
	_m.Called(language, name)
}
