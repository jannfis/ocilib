// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	distribution "github.com/distribution/distribution/v3"
	mock "github.com/stretchr/testify/mock"
)

// TagService is an autogenerated mock type for the TagService type
type TagService struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx
func (_m *TagService) All(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for All")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, tag
func (_m *TagService) Get(ctx context.Context, tag string) (distribution.Descriptor, error) {
	ret := _m.Called(ctx, tag)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 distribution.Descriptor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (distribution.Descriptor, error)); ok {
		return rf(ctx, tag)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) distribution.Descriptor); ok {
		r0 = rf(ctx, tag)
	} else {
		r0 = ret.Get(0).(distribution.Descriptor)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Lookup provides a mock function with given fields: ctx, digest
func (_m *TagService) Lookup(ctx context.Context, digest distribution.Descriptor) ([]string, error) {
	ret := _m.Called(ctx, digest)

	if len(ret) == 0 {
		panic("no return value specified for Lookup")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, distribution.Descriptor) ([]string, error)); ok {
		return rf(ctx, digest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, distribution.Descriptor) []string); ok {
		r0 = rf(ctx, digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, distribution.Descriptor) error); ok {
		r1 = rf(ctx, digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tag provides a mock function with given fields: ctx, tag, desc
func (_m *TagService) Tag(ctx context.Context, tag string, desc distribution.Descriptor) error {
	ret := _m.Called(ctx, tag, desc)

	if len(ret) == 0 {
		panic("no return value specified for Tag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, distribution.Descriptor) error); ok {
		r0 = rf(ctx, tag, desc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Untag provides a mock function with given fields: ctx, tag
func (_m *TagService) Untag(ctx context.Context, tag string) error {
	ret := _m.Called(ctx, tag)

	if len(ret) == 0 {
		panic("no return value specified for Untag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTagService creates a new instance of TagService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTagService(t interface {
	mock.TestingT
	Cleanup(func())
}) *TagService {
	mock := &TagService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
