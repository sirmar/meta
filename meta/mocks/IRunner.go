// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// IRunner is an autogenerated mock type for the IRunner type
type IRunner struct {
	mock.Mock
}

// Run provides a mock function with given fields: cmd, args
func (_m *IRunner) Run(cmd string, args []string) {
	_m.Called(cmd, args)
}