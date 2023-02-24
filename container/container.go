package container

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Concrete func(args ...any) (any, error)

// The container's bindingList.
var bindingList map[tKey]Concrete
var sharedList map[tKey]bool
var instanceList map[tKey]any
var resolvedList map[tKey]bool
var aliasList map[tKey]tKey

var di *dig.Container

func Initialize() {
	bindingList = map[tKey]Concrete{}
	sharedList = map[tKey]bool{}
	instanceList = map[tKey]any{}
	resolvedList = map[tKey]bool{}
	aliasList = map[tKey]tKey{}
	di = dig.New()
}

func getAlias(key tKey) tKey {
	val, ok := aliasList[key]
	if ok {
		return getAlias(val)
	}
	return key
}

// Bound Determine if the given abstract type has been bound.
func Bound[VT VTConstraint](abstract VT) bool {
	key := toTKey[VT](abstract)

	if _, ok := aliasList[key]; ok {
		return ok
	}

	if _, ok := bindingList[key]; ok {
		return ok
	}

	if _, ok := instanceList[key]; ok {
		return ok
	}

	return false
}

func Has[VT VTConstraint](abstract VT) bool {
	return Bound(abstract)
}

func Alias[VT1 VTConstraint, VT2 VTConstraint](abstract VT1, source VT2) {
	key1 := toTKey[VT1](abstract)
	key2 := toTKey[VT2](source)
	aliasList[key1] = key2
}

func Bind[VT VTConstraint](abstract VT, concrete Concrete, singleton bool) {
	key := toTKey[VT](abstract)
	bindingList[key] = concrete
	sharedList[key] = singleton

	// delete the alias if you explicitly bind
	delete(aliasList, key)

	// register DI provide if abstract is a non-string type
	if singleton && key.isTypeKey {
		di.Provide(func() (VT, error) {
			return resolve[VT](key)
		})
	}

	// If the abstract type was already resolved in this container we'll fire the
	// rebound listener so that any objects which have already gotten resolved
	// can have their copy of the object updated via the listener callbacks.
	if resolved(key) {
		rebound(key)
	}
}

func Singleton[VT VTConstraint](abstract VT, concrete Concrete) {
	Bind(abstract, concrete, true)
}

// Instance Register an existing instance as shared in the container.
func Instance[VT VTConstraint](abstract VT, instance any) any {
	key := toTKey[VT](abstract)
	key = getAlias(key)

	// We'll check to determine if this type has been bound before, and if it has
	// we will fire the rebound callbacks registered with the container, and it
	// can be updated with consuming classes that have gotten resolved here.
	instanceList[key] = instance
	sharedList[key] = true

	return instance
}

func resolved(key tKey) bool {
	if _, ok := resolvedList[key]; ok {
		return true
	}
	_, ok := instanceList[key]
	return ok
}

func Resolved[VT VTConstraint](abstract VT) bool {
	key := toTKey[VT](abstract)
	key = getAlias(key)

	return resolved(key)
}

func IsShared[VT VTConstraint](abstract VT) bool {
	key := toTKey[VT](abstract)
	key = getAlias(key)

	if _, ok := instanceList[key]; ok {
		return ok
	}

	_, ok := bindingList[key]
	singleton := sharedList[key]
	return ok && singleton
}

// resolve the given type from the container.
func resolve[T any](key tKey, args ...any) (T, error) {
	var obj any
	var ok bool
	var err error
	var instance T

	// not instanced
	if obj, ok = instanceList[key]; !ok {
		concrete, ok := bindingList[key]
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
		// the instanceList in "memory" so we can return it later without creating an
		// entirely new instance of an object on each subsequent request for it.
		if sharedList[key] {
			instanceList[key] = obj
		}

		resolvedList[key] = true
	}

	if instance, ok = obj.(T); ok {
		return instance, nil
	}

	return instance, errors.Errorf("[Container]: expected %s key to be type %T but got %T", key.String(), new(T), obj)
}

// Make Resolve the given type from the container.
func Make[T any, VT VTConstraint](abstract VT, args ...any) (T, error) {
	key := toTKey[VT](abstract)
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
