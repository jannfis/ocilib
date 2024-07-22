// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	digest "github.com/opencontainers/go-digest"

	mock "github.com/stretchr/testify/mock"
)

// TagManifestsProvider is an autogenerated mock type for the TagManifestsProvider type
type TagManifestsProvider struct {
	mock.Mock
}

// ManifestDigests provides a mock function with given fields: ctx, tag
func (_m *TagManifestsProvider) ManifestDigests(ctx context.Context, tag string) ([]digest.Digest, error) {
	ret := _m.Called(ctx, tag)

	if len(ret) == 0 {
		panic("no return value specified for ManifestDigests")
	}

	var r0 []digest.Digest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]digest.Digest, error)); ok {
		return rf(ctx, tag)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []digest.Digest); ok {
		r0 = rf(ctx, tag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]digest.Digest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTagManifestsProvider creates a new instance of TagManifestsProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTagManifestsProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *TagManifestsProvider {
	mock := &TagManifestsProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
