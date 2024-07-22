// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	distribution "github.com/distribution/distribution/v3"
	mock "github.com/stretchr/testify/mock"
)

// BlobDescriptorServiceFactory is an autogenerated mock type for the BlobDescriptorServiceFactory type
type BlobDescriptorServiceFactory struct {
	mock.Mock
}

// BlobAccessController provides a mock function with given fields: svc
func (_m *BlobDescriptorServiceFactory) BlobAccessController(svc distribution.BlobDescriptorService) distribution.BlobDescriptorService {
	ret := _m.Called(svc)

	if len(ret) == 0 {
		panic("no return value specified for BlobAccessController")
	}

	var r0 distribution.BlobDescriptorService
	if rf, ok := ret.Get(0).(func(distribution.BlobDescriptorService) distribution.BlobDescriptorService); ok {
		r0 = rf(svc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(distribution.BlobDescriptorService)
		}
	}

	return r0
}

// NewBlobDescriptorServiceFactory creates a new instance of BlobDescriptorServiceFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlobDescriptorServiceFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlobDescriptorServiceFactory {
	mock := &BlobDescriptorServiceFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
