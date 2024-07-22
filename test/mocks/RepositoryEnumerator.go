// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RepositoryEnumerator is an autogenerated mock type for the RepositoryEnumerator type
type RepositoryEnumerator struct {
	mock.Mock
}

// Enumerate provides a mock function with given fields: ctx, ingester
func (_m *RepositoryEnumerator) Enumerate(ctx context.Context, ingester func(string) error) error {
	ret := _m.Called(ctx, ingester)

	if len(ret) == 0 {
		panic("no return value specified for Enumerate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(string) error) error); ok {
		r0 = rf(ctx, ingester)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepositoryEnumerator creates a new instance of RepositoryEnumerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositoryEnumerator(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoryEnumerator {
	mock := &RepositoryEnumerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
