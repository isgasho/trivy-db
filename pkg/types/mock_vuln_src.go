// Code generated by mockery v1.0.0. DO NOT EDIT.

package types

import mock "github.com/stretchr/testify/mock"

// MockVulnSrc is an autogenerated mock type for the VulnSrc type
type MockVulnSrc struct {
	mock.Mock
}

// Get provides a mock function with given fields: release, pkgName
func (_m *MockVulnSrc) Get(release string, pkgName string) ([]Advisory, error) {
	ret := _m.Called(release, pkgName)

	var r0 []Advisory
	if rf, ok := ret.Get(0).(func(string, string) []Advisory); ok {
		r0 = rf(release, pkgName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Advisory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(release, pkgName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: dir
func (_m *MockVulnSrc) Update(dir string) error {
	ret := _m.Called(dir)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(dir)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
