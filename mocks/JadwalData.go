// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	jadwal "Gurumu/features/jadwal"

	mock "github.com/stretchr/testify/mock"
)

// JadwalData is an autogenerated mock type for the JadwalData type
type JadwalData struct {
	mock.Mock
}

// Add provides a mock function with given fields: guruID, newJadwal
func (_m *JadwalData) Add(guruID uint, newJadwal jadwal.Core) (jadwal.Core, error) {
	ret := _m.Called(guruID, newJadwal)

	var r0 jadwal.Core
	if rf, ok := ret.Get(0).(func(uint, jadwal.Core) jadwal.Core); ok {
		r0 = rf(guruID, newJadwal)
	} else {
		r0 = ret.Get(0).(jadwal.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, jadwal.Core) error); ok {
		r1 = rf(guruID, newJadwal)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJadwal provides a mock function with given fields: guruID
func (_m *JadwalData) GetJadwal(guruID uint) ([]jadwal.Core, error) {
	ret := _m.Called(guruID)

	var r0 []jadwal.Core
	if rf, ok := ret.Get(0).(func(uint) []jadwal.Core); ok {
		r0 = rf(guruID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]jadwal.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(guruID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewJadwalData interface {
	mock.TestingT
	Cleanup(func())
}

// NewJadwalData creates a new instance of JadwalData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJadwalData(t mockConstructorTestingTNewJadwalData) *JadwalData {
	mock := &JadwalData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
