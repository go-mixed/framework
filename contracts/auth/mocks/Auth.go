// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	auth "gopkg.in/go-mixed/framework.v1/contracts/auth"
	http "gopkg.in/go-mixed/framework.v1/contracts/http"

	mock "github.com/stretchr/testify/mock"
)

// Auth is an autogenerated mock type for the Auth type
type Auth struct {
	mock.Mock
}

// Guard provides a mock function with given fields: name
func (_m *Auth) Guard(name string) auth.IAuth {
	ret := _m.Called(name)

	var r0 auth.IAuth
	if rf, ok := ret.Get(0).(func(string) auth.IAuth); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(auth.IAuth)
		}
	}

	return r0
}

// Login provides a mock function with given fields: ctx, user
func (_m *Auth) Login(ctx http.Context, user interface{}) (string, error) {
	ret := _m.Called(ctx, user)

	var r0 string
	if rf, ok := ret.Get(0).(func(http.Context, interface{}) string); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(http.Context, interface{}) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUsingID provides a mock function with given fields: ctx, id
func (_m *Auth) LoginUsingID(ctx http.Context, id interface{}) (string, error) {
	ret := _m.Called(ctx, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(http.Context, interface{}) string); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(http.Context, interface{}) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx
func (_m *Auth) Logout(ctx http.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(http.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Parse provides a mock function with given fields: ctx, token
func (_m *Auth) Parse(ctx http.Context, token string) error {
	ret := _m.Called(ctx, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(http.Context, string) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Refresh provides a mock function with given fields: ctx
func (_m *Auth) Refresh(ctx http.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(http.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(http.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User provides a mock function with given fields: ctx, user
func (_m *Auth) User(ctx http.Context, user interface{}) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(http.Context, interface{}) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAuth interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuth creates a new instance of Auth. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuth(t mockConstructorTestingTNewAuth) *Auth {
	mock := &Auth{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
