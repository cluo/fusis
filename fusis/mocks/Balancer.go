package mocks

import fusis "github.com/luizbafilho/fusis/fusis"
import mock "github.com/stretchr/testify/mock"
import types "github.com/luizbafilho/fusis/types"

// Balancer is an autogenerated mock type for the Balancer type
type Balancer struct {
	mock.Mock
}

// AddCheck provides a mock function with given fields: check
func (_m *Balancer) AddCheck(check types.CheckSpec) error {
	ret := _m.Called(check)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.CheckSpec) error); ok {
		r0 = rf(check)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddDestination provides a mock function with given fields: _a0, _a1
func (_m *Balancer) AddDestination(_a0 *types.Service, _a1 *types.Destination) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Service, *types.Destination) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddService provides a mock function with given fields: _a0
func (_m *Balancer) AddService(_a0 *types.Service) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Service) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCheck provides a mock function with given fields: check
func (_m *Balancer) DeleteCheck(check types.CheckSpec) error {
	ret := _m.Called(check)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.CheckSpec) error); ok {
		r0 = rf(check)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDestination provides a mock function with given fields: _a0
func (_m *Balancer) DeleteDestination(_a0 *types.Destination) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Destination) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteService provides a mock function with given fields: _a0
func (_m *Balancer) DeleteService(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDestination provides a mock function with given fields: _a0
func (_m *Balancer) GetDestination(_a0 string) (*types.Destination, error) {
	ret := _m.Called(_a0)

	var r0 *types.Destination
	if rf, ok := ret.Get(0).(func(string) *types.Destination); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Destination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDestinations provides a mock function with given fields: svc
func (_m *Balancer) GetDestinations(svc *types.Service) []types.Destination {
	ret := _m.Called(svc)

	var r0 []types.Destination
	if rf, ok := ret.Get(0).(func(*types.Service) []types.Destination); ok {
		r0 = rf(svc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Destination)
		}
	}

	return r0
}

// GetService provides a mock function with given fields: _a0
func (_m *Balancer) GetService(_a0 string) (*types.Service, error) {
	ret := _m.Called(_a0)

	var r0 *types.Service
	if rf, ok := ret.Get(0).(func(string) *types.Service); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServices provides a mock function with given fields:
func (_m *Balancer) GetServices() []types.Service {
	ret := _m.Called()

	var r0 []types.Service
	if rf, ok := ret.Get(0).(func() []types.Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Service)
		}
	}

	return r0
}

// IsLeader provides a mock function with given fields:
func (_m *Balancer) IsLeader() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Shutdown provides a mock function with given fields:
func (_m *Balancer) Shutdown() {
	_m.Called()
}

var _ fusis.Balancer = (*Balancer)(nil)
