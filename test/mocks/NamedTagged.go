// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NamedTagged is an autogenerated mock type for the NamedTagged type
type NamedTagged struct {
	mock.Mock
}

// Name provides a mock function with given fields:
func (_m *NamedTagged) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *NamedTagged) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Tag provides a mock function with given fields:
func (_m *NamedTagged) Tag() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Tag")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewNamedTagged creates a new instance of NamedTagged. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNamedTagged(t interface {
	mock.TestingT
	Cleanup(func())
}) *NamedTagged {
	mock := &NamedTagged{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
