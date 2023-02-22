// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	log "gopkg.in/go-mixed/framework/contracts/log"
)

// Log is an autogenerated mock type for the Log type
type Log struct {
	mock.Mock
}

// Debug provides a mock function with given fields: args
func (_m *Log) Debug(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Debugf provides a mock function with given fields: format, args
func (_m *Log) Debugf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: args
func (_m *Log) Error(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Errorf provides a mock function with given fields: format, args
func (_m *Log) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatal provides a mock function with given fields: args
func (_m *Log) Fatal(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatalf provides a mock function with given fields: format, args
func (_m *Log) Fatalf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: args
func (_m *Log) Info(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Infof provides a mock function with given fields: format, args
func (_m *Log) Infof(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Panic provides a mock function with given fields: args
func (_m *Log) Panic(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Panicf provides a mock function with given fields: format, args
func (_m *Log) Panicf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warning provides a mock function with given fields: args
func (_m *Log) Warning(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warningf provides a mock function with given fields: format, args
func (_m *Log) Warningf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// WithContext provides a mock function with given fields: ctx
func (_m *Log) WithContext(ctx context.Context) log.Writer {
	ret := _m.Called(ctx)

	var r0 log.Writer
	if rf, ok := ret.Get(0).(func(context.Context) log.Writer); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Writer)
		}
	}

	return r0
}

type mockConstructorTestingTNewLog interface {
	mock.TestingT
	Cleanup(func())
}

// NewLog creates a new instance of Log. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLog(t mockConstructorTestingTNewLog) *Log {
	mock := &Log{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
