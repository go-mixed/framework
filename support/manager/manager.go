package manager

import (
	"github.com/pkg/errors"
	"gopkg.in/go-mixed/framework/contracts/manager"
)

type Manager[T any] struct {
	drivers           map[string]T
	customCreators    map[string]manager.Concrete[T]
	defaultDriverName func() string
}

func MakeManager[T any](defaultDriverName func() string) Manager[T] {
	return Manager[T]{
		drivers:           map[string]T{},
		customCreators:    map[string]manager.Concrete[T]{},
		defaultDriverName: defaultDriverName,
	}
}

// Extend Register a custom driver creator Closure.
func (m *Manager[T]) Extend(name string, concrete manager.Concrete[T]) *Manager[T] {
	m.customCreators[name] = concrete
	return m
}

func (m *Manager[T]) HasDriver(name string) bool {
	if _, ok := m.drivers[name]; ok {
		return ok
	}
	_, ok := m.customCreators[name]
	return ok
}

// Driver Get a driver instance.
func (m *Manager[T]) Driver(name string) (T, error) {
	var err error
	instance, ok := m.drivers[name]
	if !ok {
		if concrete, ok := m.customCreators[name]; ok {
			if instance, err = concrete(name); err != nil {
				return nil, err
			}

			m.drivers[name] = instance
		} else {
			return nil, errors.Errorf("driver %s is not exists", name)
		}
	}

	if instance == nil {
		return nil, errors.Errorf("driver %s cannot make a valid instance", name)
	}

	return instance, nil
}

func (m *Manager[T]) MustDriver(name string) T {
	instance, err := m.Driver(name)
	if err != nil {
		panic(err.Error())
	}

	return instance
}

func (m *Manager[T]) RemoveDriver(name string) {
	delete(m.drivers, name)
	delete(m.customCreators, name)
}

func (m *Manager[T]) DefaultDriver() (T, error) {
	return m.Driver(m.defaultDriverName())
}
