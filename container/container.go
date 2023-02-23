package container

import (
	"github.com/eddieowens/axon"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Concrete func(args ...any) (any, error)

// The container's bindings.
var bindings map[tKey]Concrete
var shared map[tKey]bool
var instances map[tKey]any
var resolved map[tKey]bool
var alias map[tKey]tKey

var di *dig.Container

type VTConstraint interface {
	any
}

func Initialize() {
	bindings = map[tKey]Concrete{}
	shared = map[tKey]bool{}
	instances = map[tKey]any{}
	resolved = map[tKey]bool{}
	alias = map[tKey]tKey{}
	di = dig.New()
}

func getAlias(key tKey) tKey {
	val, ok := alias[key]
	if ok {
		return getAlias(val)
	}
	return key
}

// Bound Determine if the given abstract type has been bound.
func Bound[VT VTConstraint](abstract VT) bool {
	key := toTKey(abstract)

	if _, ok := alias[key]; ok {
		return ok
	}

	if _, ok := bindings[key]; ok {
		return ok
	}

	if _, ok := instances[key]; ok {
		return ok
	}

	return false
}

func Has[VT VTConstraint](abstract VT) bool {
	return Bound(abstract)
}

func Alias[VT VTConstraint](abstract VT, source any) {
	key1 := toTKey(abstract)
	key2 := toTKey(source)
	alias[key1] = key2
}

func Bind[VT VTConstraint](abstract VT, concrete Concrete, singleton bool) {
	key := toTKey(abstract)
	bindings[key] = concrete
	shared[key] = singleton

	// delete the alias if you explicitly bind
	delete(alias, key)

	// register DI provide if abstract is a non-string type
	if singleton && key.isTypeKey {
		di.Provide(func() (VT, error) {
			return resolve[VT](key)
		})
	}

	// If the abstract type was already resolved in this container we'll fire the
	// rebound listener so that any objects which have already gotten resolved
	// can have their copy of the object updated via the listener callbacks.
	if Resolved(key) {
		rebound(key)
	}
}

func Singleton[VT VTConstraint](abstract VT, concrete Concrete) {
	Bind(abstract, concrete, true)
}

// Instance Register an existing instance as shared in the container.
func Instance[VT VTConstraint](abstract VT, instance any) any {
	key := toTKey(abstract)
	key = getAlias(key)

	// We'll check to determine if this type has been bound before, and if it has
	// we will fire the rebound callbacks registered with the container, and it
	// can be updated with consuming classes that have gotten resolved here.
	instances[key] = instance
	shared[key] = true

	return instance
}

func Resolved[VT VTConstraint](abstract VT) bool {
	key := toTKey(abstract)
	key = getAlias(key)

	if _, ok := resolved[key]; ok {
		return true
	}
	_, ok := instances[key]
	return ok
}

func IsShared[VT VTConstraint](abstract VT) bool {
	key := toTKey(abstract)
	key = getAlias(key)

	if _, ok := instances[key]; ok {
		return ok
	}

	_, ok := bindings[key]
	singleton := shared[key]
	return ok && singleton
}

// resolve the given type from the container.
func resolve[T any](key tKey, args ...any) (T, error) {
	var obj any
	var ok bool
	var err error

	instance := obj.(T)

	// not instanced
	if obj, ok = instances[key]; !ok {
		concrete, ok := bindings[key]
		if !ok {
			return instance, errors.Errorf("abstract \"%s\" is not exists in the container", key.String())
		}

		if concrete == nil {
			return instance, errors.Errorf("the concrete \"%s\" is invalid in the container", key.String())
		}

		obj, err = concrete(args...)
		if err != nil {
			return instance, err
		}

		// If the requested type is registered as a singleton we'll want to cache off
		// the instances in "memory" so we can return it later without creating an
		// entirely new instance of an object on each subsequent request for it.
		if shared[key] {
			instances[key] = obj
		}

		resolved[key] = true
	}

	if instance, ok = obj.(T); ok {
		return instance, nil
	}

	return instance, errors.Errorf("%s: expected %s key to be type %T but got %T", axon.ErrInvalidType.Error(), key.String(), new(T), obj)
}

// Make Resolve the given type from the container.
func Make[T any, VT VTConstraint](abstract VT, args ...any) (T, error) {
	key := toTKey(abstract)
	key = getAlias(key)
	return resolve[T](key, args...)
}

func MustMake[T any, VT VTConstraint](abstract VT, args ...any) T {
	instance, err := Make[T](abstract, args...)
	if err != nil {
		panic(err)
	}

	return instance
}

func rebound(key tKey) {
	_, _ = resolve[any](key)
}

// Invoke a func via DI(IoC)
func Invoke(fn any) error {
	return di.Invoke(fn)
}
