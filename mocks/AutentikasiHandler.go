// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// AutentikasiHandler is an autogenerated mock type for the AutentikasiHandler type
type AutentikasiHandler struct {
	mock.Mock
}

// Login provides a mock function with given fields:
func (_m *AutentikasiHandler) Login() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

type mockConstructorTestingTNewAutentikasiHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewAutentikasiHandler creates a new instance of AutentikasiHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAutentikasiHandler(t mockConstructorTestingTNewAutentikasiHandler) *AutentikasiHandler {
	mock := &AutentikasiHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
