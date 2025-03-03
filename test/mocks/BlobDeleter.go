// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	digest "github.com/opencontainers/go-digest"

	mock "github.com/stretchr/testify/mock"
)

// BlobDeleter is an autogenerated mock type for the BlobDeleter type
type BlobDeleter struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, dgst
func (_m *BlobDeleter) Delete(ctx context.Context, dgst digest.Digest) error {
	ret := _m.Called(ctx, dgst)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, digest.Digest) error); ok {
		r0 = rf(ctx, dgst)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBlobDeleter creates a new instance of BlobDeleter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlobDeleter(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlobDeleter {
	mock := &BlobDeleter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
