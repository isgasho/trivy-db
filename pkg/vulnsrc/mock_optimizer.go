// Code generated by mockery v1.0.0. DO NOT EDIT.

package vulnsrc

import mock "github.com/stretchr/testify/mock"

// MockOptimizer is an autogenerated mock type for the Optimizer type
type MockOptimizer struct {
	mock.Mock
}

// Optimize provides a mock function with given fields:
func (_m *MockOptimizer) Optimize() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
