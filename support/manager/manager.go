package manager

import (
	"github.com/pkg/errors"
	"gopkg.in/go-mixed/framework.v1/contracts/manager"
	"sync"
)

type Manager[T any] struct {
	instances         sync.Map
	customCreators    map[string]manager.Concrete[T]
	defaultDriverName func() string
	resolve           func(string) (T, error)
	mu                sync.Mutex
}

var _ manager.IManager[any] = (*Manager[any])(nil)

func MakeManager[T any](defaultDriverName func() string, resolve func(string) (T, error)) Manager[T] {
	return Manager[T]{
		instances:         sync.Map{},
		customCreators:    map[string]manager.Concrete[T]{},
		defaultDriverName: defaultDriverName,
		resolve:           resolve,
		mu:                sync.Mutex{},
	}
}

// Extend Register a custom driver creator Closure.
func (m *Manager[T]) Extend(driverName string, concrete manager.Concrete[T]) *Manager[T] {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.customCreators[driverName] = concrete
	return m
}

func (m *Manager[T]) Resolved(name string) bool {
	_, ok := m.instances.Load(name)
	return ok
}

func (m *Manager[T]) HasCustomCreator(name string) bool {
	_, ok := m.customCreators[name]
	return ok
}

// Driver Get a driver instance.
func (m *Manager[T]) Driver(name string) (T, error) {
	var err error
	var instance T
	if name == "" {
		name = m.DefaultDriverName()
	}

	obj, ok := m.instances.Load(name)
	if !ok {
		instance, err = m.resolve(name)
		if err != nil {
			return instance, err
		}

		m.instances.Store(name, instance)

		return instance, nil
	} else if instance, ok = obj.(T); ok {
		return instance, nil
	}

	return instance, errors.Errorf("the instance of driver %s cannot cast %T to type %T", name, instance, new(T))
}

func (m *Manager[T]) CallCustomCreator(creatorName, driverName string, args ...any) (T, error) {
	var instance T
	var err error
	if concrete, ok := m.customCreators[creatorName]; ok {
		instance, err = concrete(driverName, args...)
	} else {
		err = errors.Errorf("driver \"%s.%s\" is not exists", driverName, creatorName)
	}

	return instance, err
}

func (m *Manager[T]) MustDriver(name string) T {
	instance, err := m.Driver(name)
	if err != nil {
		panic(err.Error())
	}

	return instance
}

func (m *Manager[T]) Remove(name string) {
	m.instances.Delete(name)
}

func (m *Manager[T]) RemoveCustomCreator(name string) {
	delete(m.customCreators, name)
}

func (m *Manager[T]) DefaultDriver() (T, error) {
	return m.Driver(m.defaultDriverName())
}

func (m *Manager[T]) MustDefaultDriver() T {
	driver, err := m.DefaultDriver()
	if err != nil {
		panic(err)
	}

	return driver
}

func (m *Manager[T]) DefaultDriverName() string {
	return m.defaultDriverName()
}
