// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	filesystem "gopkg.in/go-mixed/framework.v1/contracts/filesystem"

	time "time"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// AllDirectories provides a mock function with given fields: path
func (_m *Storage) AllDirectories(path string) ([]string, error) {
	ret := _m.Called(path)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllFiles provides a mock function with given fields: path
func (_m *Storage) AllFiles(path string) ([]string, error) {
	ret := _m.Called(path)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Copy provides a mock function with given fields: oldFile, newFile
func (_m *Storage) Copy(oldFile string, newFile string) error {
	ret := _m.Called(oldFile, newFile)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(oldFile, newFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: file
func (_m *Storage) Delete(file ...string) error {
	_va := make([]interface{}, len(file))
	for _i := range file {
		_va[_i] = file[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(file...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDirectory provides a mock function with given fields: directory
func (_m *Storage) DeleteDirectory(directory string) error {
	ret := _m.Called(directory)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(directory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Directories provides a mock function with given fields: path
func (_m *Storage) Directories(path string) ([]string, error) {
	ret := _m.Called(path)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Disk provides a mock function with given fields: disk
func (_m *Storage) Disk(disk string) filesystem.Driver {
	ret := _m.Called(disk)

	var r0 filesystem.Driver
	if rf, ok := ret.Get(0).(func(string) filesystem.Driver); ok {
		r0 = rf(disk)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(filesystem.Driver)
		}
	}

	return r0
}

// Exists provides a mock function with given fields: file
func (_m *Storage) Exists(file string) bool {
	ret := _m.Called(file)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Files provides a mock function with given fields: path
func (_m *Storage) Files(path string) ([]string, error) {
	ret := _m.Called(path)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: file
func (_m *Storage) Get(file string) (string, error) {
	ret := _m.Called(file)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeDirectory provides a mock function with given fields: directory
func (_m *Storage) MakeDirectory(directory string) error {
	ret := _m.Called(directory)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(directory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Missing provides a mock function with given fields: file
func (_m *Storage) Missing(file string) bool {
	ret := _m.Called(file)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Move provides a mock function with given fields: oldFile, newFile
func (_m *Storage) Move(oldFile string, newFile string) error {
	ret := _m.Called(oldFile, newFile)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(oldFile, newFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Path provides a mock function with given fields: file
func (_m *Storage) Path(file string) string {
	ret := _m.Called(file)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Put provides a mock function with given fields: file, content
func (_m *Storage) Put(file string, content string) error {
	ret := _m.Called(file, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(file, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutFile provides a mock function with given fields: path, source
func (_m *Storage) PutFile(path string, source filesystem.File) (string, error) {
	ret := _m.Called(path, source)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, filesystem.File) string); ok {
		r0 = rf(path, source)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, filesystem.File) error); ok {
		r1 = rf(path, source)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutFileAs provides a mock function with given fields: path, source, name
func (_m *Storage) PutFileAs(path string, source filesystem.File, name string) (string, error) {
	ret := _m.Called(path, source, name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, filesystem.File, string) string); ok {
		r0 = rf(path, source, name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, filesystem.File, string) error); ok {
		r1 = rf(path, source, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Size provides a mock function with given fields: file
func (_m *Storage) Size(file string) (int64, error) {
	ret := _m.Called(file)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TemporaryUrl provides a mock function with given fields: file, _a1
func (_m *Storage) TemporaryUrl(file string, _a1 time.Time) (string, error) {
	ret := _m.Called(file, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, time.Time) string); ok {
		r0 = rf(file, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, time.Time) error); ok {
		r1 = rf(file, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Url provides a mock function with given fields: file
func (_m *Storage) Url(file string) string {
	ret := _m.Called(file)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WithContext provides a mock function with given fields: ctx
func (_m *Storage) WithContext(ctx context.Context) filesystem.Driver {
	ret := _m.Called(ctx)

	var r0 filesystem.Driver
	if rf, ok := ret.Get(0).(func(context.Context) filesystem.Driver); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(filesystem.Driver)
		}
	}

	return r0
}

type mockConstructorTestingTNewStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorage(t mockConstructorTestingTNewStorage) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
