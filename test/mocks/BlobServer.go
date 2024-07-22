// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	digest "github.com/opencontainers/go-digest"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// BlobServer is an autogenerated mock type for the BlobServer type
type BlobServer struct {
	mock.Mock
}

// ServeBlob provides a mock function with given fields: ctx, w, r, dgst
func (_m *BlobServer) ServeBlob(ctx context.Context, w http.ResponseWriter, r *http.Request, dgst digest.Digest) error {
	ret := _m.Called(ctx, w, r, dgst)

	if len(ret) == 0 {
		panic("no return value specified for ServeBlob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, http.ResponseWriter, *http.Request, digest.Digest) error); ok {
		r0 = rf(ctx, w, r, dgst)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBlobServer creates a new instance of BlobServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlobServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlobServer {
	mock := &BlobServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
