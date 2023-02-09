// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	reservasi "Gurumu/features/reservasi"

	mock "github.com/stretchr/testify/mock"
)

// ReservasiData is an autogenerated mock type for the ReservasiData type
type ReservasiData struct {
	mock.Mock
}

// Add provides a mock function with given fields: siswaID, newReservasi
func (_m *ReservasiData) Add(siswaID uint, newReservasi reservasi.Core) (reservasi.Core, error) {
	ret := _m.Called(siswaID, newReservasi)

	var r0 reservasi.Core
	if rf, ok := ret.Get(0).(func(uint, reservasi.Core) reservasi.Core); ok {
		r0 = rf(siswaID, newReservasi)
	} else {
		r0 = ret.Get(0).(reservasi.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, reservasi.Core) error); ok {
		r1 = rf(siswaID, newReservasi)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mysession provides a mock function with given fields: userID, role, reservasiStatus
func (_m *ReservasiData) Mysession(userID uint, role string, reservasiStatus string) ([]reservasi.Core, error) {
	ret := _m.Called(userID, role, reservasiStatus)

	var r0 []reservasi.Core
	if rf, ok := ret.Get(0).(func(uint, string, string) []reservasi.Core); ok {
		r0 = rf(userID, role, reservasiStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reservasi.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string, string) error); ok {
		r1 = rf(userID, role, reservasiStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewReservasiData interface {
	mock.TestingT
	Cleanup(func())
}

// NewReservasiData creates a new instance of ReservasiData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReservasiData(t mockConstructorTestingTNewReservasiData) *ReservasiData {
	mock := &ReservasiData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
