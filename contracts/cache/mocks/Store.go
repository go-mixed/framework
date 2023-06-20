// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	contractscache "gopkg.in/go-mixed/framework.v1/contracts/cache"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

var _ contractscache.IStore = (*Store)(nil)

func (_m *Store) Store(storeName string) contractscache.IStore {
	return &Store{}
}

// Add provides a mock function with given fields: key, value, sec
func (_m *Store) Add(key string, value interface{}, sec time.Duration) bool {
	ret := _m.Called(key, value, sec)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, interface{}, time.Duration) bool); ok {
		r0 = rf(key, value, sec)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Flush provides a mock function with given fields:
func (_m *Store) Flush() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Forever provides a mock function with given fields: key, value
func (_m *Store) Forever(key string, value interface{}) bool {
	ret := _m.Called(key, value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, interface{}) bool); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Forget provides a mock function with given fields: key
func (_m *Store) Forget(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Get provides a mock function with given fields: key, def
func (_m *Store) Get(key string, def interface{}) interface{} {
	ret := _m.Called(key, def)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, interface{}) interface{}); ok {
		r0 = rf(key, def)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GetBool provides a mock function with given fields: key, def
func (_m *Store) GetBool(key string, def bool) bool {
	ret := _m.Called(key, def)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, bool) bool); ok {
		r0 = rf(key, def)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetInt provides a mock function with given fields: key, def
func (_m *Store) GetInt(key string, def int) int {
	ret := _m.Called(key, def)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, int) int); ok {
		r0 = rf(key, def)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetString provides a mock function with given fields: key, def
func (_m *Store) GetString(key string, def string) string {
	ret := _m.Called(key, def)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(key, def)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Has provides a mock function with given fields: key
func (_m *Store) Has(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Pull provides a mock function with given fields: key, def
func (_m *Store) Pull(key string, def interface{}) interface{} {
	ret := _m.Called(key, def)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, interface{}) interface{}); ok {
		r0 = rf(key, def)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Put provides a mock function with given fields: key, value, sec
func (_m *Store) Put(key string, value interface{}, sec time.Duration) error {
	ret := _m.Called(key, value, sec)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, time.Duration) error); ok {
		r0 = rf(key, value, sec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remember provides a mock function with given fields: key, ttl, callback
func (_m *Store) Remember(key string, ttl time.Duration, callback func() interface{}) (interface{}, error) {
	ret := _m.Called(key, ttl, callback)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, time.Duration, func() interface{}) interface{}); ok {
		r0 = rf(key, ttl, callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, time.Duration, func() interface{}) error); ok {
		r1 = rf(key, ttl, callback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RememberForever provides a mock function with given fields: key, callback
func (_m *Store) RememberForever(key string, callback func() interface{}) (interface{}, error) {
	ret := _m.Called(key, callback)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, func() interface{}) interface{}); ok {
		r0 = rf(key, callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, func() interface{}) error); ok {
		r1 = rf(key, callback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithContext provides a mock function with given fields: ctx
func (_m *Store) WithContext(ctx context.Context) contractscache.IStore {
	ret := _m.Called(ctx)

	var r0 contractscache.IStore
	if rf, ok := ret.Get(0).(func(context.Context) contractscache.IStore); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(contractscache.IStore)
		}
	}

	return r0
}

// Put provides a mock function with given fields: key, value, sec
func (_m *Store) ClearPrefix(delPrefix string) error {
	ret := _m.Called(delPrefix)

	var r0 error
	if rf, ok := ret.Get(0).(func(delPrefix string) error); ok {
		r0 = rf(delPrefix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStore(t mockConstructorTestingTNewStore) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
