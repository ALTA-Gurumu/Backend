// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	jadwal "Gurumu/features/jadwal"

	mock "github.com/stretchr/testify/mock"
)

// JadwalService is an autogenerated mock type for the JadwalService type
type JadwalService struct {
	mock.Mock
}

// Add provides a mock function with given fields: token, newJadwal
func (_m *JadwalService) Add(token interface{}, newJadwal jadwal.Core) (jadwal.Core, error) {
	ret := _m.Called(token, newJadwal)

	var r0 jadwal.Core
	if rf, ok := ret.Get(0).(func(interface{}, jadwal.Core) jadwal.Core); ok {
		r0 = rf(token, newJadwal)
	} else {
		r0 = ret.Get(0).(jadwal.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, jadwal.Core) error); ok {
		r1 = rf(token, newJadwal)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJadwal provides a mock function with given fields: token
func (_m *JadwalService) GetJadwal(token interface{}) ([]jadwal.Core, error) {
	ret := _m.Called(token)

	var r0 []jadwal.Core
	if rf, ok := ret.Get(0).(func(interface{}) []jadwal.Core); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]jadwal.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewJadwalService interface {
	mock.TestingT
	Cleanup(func())
}

// NewJadwalService creates a new instance of JadwalService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJadwalService(t mockConstructorTestingTNewJadwalService) *JadwalService {
	mock := &JadwalService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
