// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// ReservasiHandler is an autogenerated mock type for the ReservasiHandler type
type ReservasiHandler struct {
	mock.Mock
}

// Add provides a mock function with given fields:
func (_m *ReservasiHandler) Add() echo.HandlerFunc {
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

// Callback provides a mock function with given fields:
func (_m *ReservasiHandler) Callback() echo.HandlerFunc {
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

// Mysession provides a mock function with given fields:
func (_m *ReservasiHandler) Mysession() echo.HandlerFunc {
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

// NotificationTransactionStatus provides a mock function with given fields:
func (_m *ReservasiHandler) NotificationTransactionStatus() echo.HandlerFunc {
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

// UpdateStatus provides a mock function with given fields:
func (_m *ReservasiHandler) UpdateStatus() echo.HandlerFunc {
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

type mockConstructorTestingTNewReservasiHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewReservasiHandler creates a new instance of ReservasiHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReservasiHandler(t mockConstructorTestingTNewReservasiHandler) *ReservasiHandler {
	mock := &ReservasiHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
