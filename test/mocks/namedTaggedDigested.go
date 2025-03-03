// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	digest "github.com/opencontainers/go-digest"
	mock "github.com/stretchr/testify/mock"
)

// namedTaggedDigested is an autogenerated mock type for the namedTaggedDigested type
type namedTaggedDigested struct {
	mock.Mock
}

// Digest provides a mock function with given fields:
func (_m *namedTaggedDigested) Digest() digest.Digest {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Digest")
	}

	var r0 digest.Digest
	if rf, ok := ret.Get(0).(func() digest.Digest); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(digest.Digest)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *namedTaggedDigested) Name() string {
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
func (_m *namedTaggedDigested) String() string {
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
func (_m *namedTaggedDigested) Tag() string {
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

// newNamedTaggedDigested creates a new instance of namedTaggedDigested. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newNamedTaggedDigested(t interface {
	mock.TestingT
	Cleanup(func())
}) *namedTaggedDigested {
	mock := &namedTaggedDigested{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
