// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	distribution "github.com/distribution/distribution/v3"
	mock "github.com/stretchr/testify/mock"

	reference "github.com/distribution/distribution/v3/reference"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Blobs provides a mock function with given fields: ctx
func (_m *Repository) Blobs(ctx context.Context) distribution.BlobStore {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Blobs")
	}

	var r0 distribution.BlobStore
	if rf, ok := ret.Get(0).(func(context.Context) distribution.BlobStore); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(distribution.BlobStore)
		}
	}

	return r0
}

// Manifests provides a mock function with given fields: ctx, options
func (_m *Repository) Manifests(ctx context.Context, options ...distribution.ManifestServiceOption) (distribution.ManifestService, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Manifests")
	}

	var r0 distribution.ManifestService
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...distribution.ManifestServiceOption) (distribution.ManifestService, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...distribution.ManifestServiceOption) distribution.ManifestService); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(distribution.ManifestService)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...distribution.ManifestServiceOption) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Named provides a mock function with given fields:
func (_m *Repository) Named() reference.Named {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Named")
	}

	var r0 reference.Named
	if rf, ok := ret.Get(0).(func() reference.Named); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reference.Named)
		}
	}

	return r0
}

// Tags provides a mock function with given fields: ctx
func (_m *Repository) Tags(ctx context.Context) distribution.TagService {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Tags")
	}

	var r0 distribution.TagService
	if rf, ok := ret.Get(0).(func(context.Context) distribution.TagService); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(distribution.TagService)
		}
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
