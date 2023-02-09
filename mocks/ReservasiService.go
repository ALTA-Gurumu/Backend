// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	reservasi "Gurumu/features/reservasi"

	mock "github.com/stretchr/testify/mock"
)

// ReservasiService is an autogenerated mock type for the ReservasiService type
type ReservasiService struct {
	mock.Mock
}

// Add provides a mock function with given fields: token, newReservasi
func (_m *ReservasiService) Add(token interface{}, newReservasi reservasi.Core) (reservasi.Core, error) {
	ret := _m.Called(token, newReservasi)

	var r0 reservasi.Core
	if rf, ok := ret.Get(0).(func(interface{}, reservasi.Core) reservasi.Core); ok {
		r0 = rf(token, newReservasi)
	} else {
		r0 = ret.Get(0).(reservasi.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, reservasi.Core) error); ok {
		r1 = rf(token, newReservasi)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mysession provides a mock function with given fields: token, role, reservasiStatus
func (_m *ReservasiService) Mysession(token interface{}, role string, reservasiStatus string) ([]reservasi.Core, error) {
	ret := _m.Called(token, role, reservasiStatus)

	var r0 []reservasi.Core
	if rf, ok := ret.Get(0).(func(interface{}, string, string) []reservasi.Core); ok {
		r0 = rf(token, role, reservasiStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reservasi.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, string, string) error); ok {
		r1 = rf(token, role, reservasiStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mysession provides a mock function with given fields: token, role, reservasiStatus
func (_m *ReservasiService) Mysession(token interface{}, role string, reservasiStatus string) ([]reservasi.Core, error) {
	ret := _m.Called(token, role, reservasiStatus)

	var r0 []reservasi.Core
	if rf, ok := ret.Get(0).(func(interface{}, string, string) []reservasi.Core); ok {
		r0 = rf(token, role, reservasiStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reservasi.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, string, string) error); ok {
		r1 = rf(token, role, reservasiStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewReservasiService interface {
	mock.TestingT
	Cleanup(func())
}

// NewReservasiService creates a new instance of ReservasiService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReservasiService(t mockConstructorTestingTNewReservasiService) *ReservasiService {
	mock := &ReservasiService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
