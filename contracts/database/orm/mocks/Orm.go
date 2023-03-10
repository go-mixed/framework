// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	"gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
	ormcontracts "gopkg.in/go-mixed/framework.v1/contracts/database/orm"

	sql "database/sql"
)

// Orm is an autogenerated mock type for the Orm type
type Orm struct {
	mock.Mock
}

var _ ormcontracts.IOrm = (*Orm)(nil)

// Connection provides a mock function with given fields: name
func (_m *Orm) Connection(name string) ormcontracts.IOrm {
	ret := _m.Called(name)

	var r0 ormcontracts.IOrm
	if rf, ok := ret.Get(0).(func(string) ormcontracts.IOrm); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ormcontracts.IOrm)
		}
	}

	return r0
}

// DB provides a mock function with given fields:
func (_m *Orm) DB() (*sql.DB, error) {
	ret := _m.Called()

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB provides a mock function with given fields:
func (_m *Orm) Gorm() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}
	return r0
}

// Query provides a mock function with given fields:
func (_m *Orm) Query() ormcontracts.DB {
	ret := _m.Called()

	var r0 ormcontracts.DB
	if rf, ok := ret.Get(0).(func() ormcontracts.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ormcontracts.DB)
		}
	}

	return r0
}

// Transaction provides a mock function with given fields: txFunc
func (_m *Orm) Transaction(txFunc func(ormcontracts.Transaction) error) error {
	ret := _m.Called(txFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(ormcontracts.Transaction) error) error); ok {
		r0 = rf(txFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithContext provides a mock function with given fields: ctx
func (_m *Orm) WithContext(ctx context.Context) ormcontracts.IOrm {
	ret := _m.Called(ctx)

	var r0 ormcontracts.IOrm
	if rf, ok := ret.Get(0).(func(context.Context) ormcontracts.IOrm); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ormcontracts.IOrm)
		}
	}

	return r0
}

type mockConstructorTestingTNewOrm interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrm creates a new instance of Orm. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrm(t mockConstructorTestingTNewOrm) *Orm {
	mock := &Orm{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
