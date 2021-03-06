// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

// AddressDeleter is an autogenerated mock type for the AddressDeleter type
type AddressDeleter struct {
	mock.Mock
}

// DeleteAddressableById provides a mock function with given fields: id
func (_m *AddressDeleter) DeleteAddressableById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAddressableById provides a mock function with given fields: id
func (_m *AddressDeleter) GetAddressableById(id string) (models.Addressable, error) {
	ret := _m.Called(id)

	var r0 models.Addressable
	if rf, ok := ret.Get(0).(func(string) models.Addressable); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Addressable)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressableByName provides a mock function with given fields: n
func (_m *AddressDeleter) GetAddressableByName(n string) (models.Addressable, error) {
	ret := _m.Called(n)

	var r0 models.Addressable
	if rf, ok := ret.Get(0).(func(string) models.Addressable); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(models.Addressable)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressables provides a mock function with given fields:
func (_m *AddressDeleter) GetAddressables() ([]models.Addressable, error) {
	ret := _m.Called()

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func() []models.Addressable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressablesByAddress provides a mock function with given fields: address
func (_m *AddressDeleter) GetAddressablesByAddress(address string) ([]models.Addressable, error) {
	ret := _m.Called(address)

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func(string) []models.Addressable); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressablesByPort provides a mock function with given fields: p
func (_m *AddressDeleter) GetAddressablesByPort(p int) ([]models.Addressable, error) {
	ret := _m.Called(p)

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func(int) []models.Addressable); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressablesByPublisher provides a mock function with given fields: p
func (_m *AddressDeleter) GetAddressablesByPublisher(p string) ([]models.Addressable, error) {
	ret := _m.Called(p)

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func(string) []models.Addressable); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressablesByTopic provides a mock function with given fields: t
func (_m *AddressDeleter) GetAddressablesByTopic(t string) ([]models.Addressable, error) {
	ret := _m.Called(t)

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func(string) []models.Addressable); ok {
		r0 = rf(t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceServicesByAddressableId provides a mock function with given fields: id
func (_m *AddressDeleter) GetDeviceServicesByAddressableId(id string) ([]models.DeviceService, error) {
	ret := _m.Called(id)

	var r0 []models.DeviceService
	if rf, ok := ret.Get(0).(func(string) []models.DeviceService); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceService)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
