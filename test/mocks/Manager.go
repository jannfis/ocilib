// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	challenge "github.com/distribution/distribution/v3/registry/client/auth/challenge"

	mock "github.com/stretchr/testify/mock"

	url "net/url"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// AddResponse provides a mock function with given fields: resp
func (_m *Manager) AddResponse(resp *http.Response) error {
	ret := _m.Called(resp)

	if len(ret) == 0 {
		panic("no return value specified for AddResponse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*http.Response) error); ok {
		r0 = rf(resp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetChallenges provides a mock function with given fields: endpoint
func (_m *Manager) GetChallenges(endpoint url.URL) ([]challenge.Challenge, error) {
	ret := _m.Called(endpoint)

	if len(ret) == 0 {
		panic("no return value specified for GetChallenges")
	}

	var r0 []challenge.Challenge
	var r1 error
	if rf, ok := ret.Get(0).(func(url.URL) ([]challenge.Challenge, error)); ok {
		return rf(endpoint)
	}
	if rf, ok := ret.Get(0).(func(url.URL) []challenge.Challenge); ok {
		r0 = rf(endpoint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]challenge.Challenge)
		}
	}

	if rf, ok := ret.Get(1).(func(url.URL) error); ok {
		r1 = rf(endpoint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
