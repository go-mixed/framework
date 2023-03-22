package container

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"sync"
)

type Concrete func(args ...any) (any, error)

type concreteContainer struct {
	concrete Concrete
	shared   bool
	instance any
	resolved bool
}

// The container's bindingList. a concurrent map, supported high-performance concurrent reading
var bindingList sync.Map

// The alias list, a normal map, use the global mutex to write, because it writes with few times
var aliasList map[tKey]tKey
var mu sync.Mutex
var di *dig.Container

func Initialize() {
	bindingList = sync.Map{}
	aliasList = map[tKey]tKey{}
	mu = sync.Mutex{}
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

	_, ok := bindingList.Load(key)
	return ok
}

func Has[VT VTConstraint](abstract VT) bool {
	return Bound(abstract)
}

func Alias[VT1 VTConstraint, VT2 VTConstraint](abstract VT1, source VT2) {
	key1 := toTKey[VT1](abstract)
	key2 := toTKey[VT2](source)
	mu.Lock()
	defer mu.Unlock()

	// register DI provide if abstract is non-string type, and the source is singleton
	if key1.isTypeKey && IsShared(source) {
		_ = registerProvider[VT1](key1)
	}
	aliasList[key1] = key2
}

func registerProvider[VT VTConstraint](key tKey) error {
	return di.Provide(func() (VT, error) {
		return resolve[VT](key)
	})
}

func Bind[VT VTConstraint](abstract VT, concrete Concrete, singleton bool) {
	key := toTKey[VT](abstract)
	_resolved := resolved(key)

	bindingList.Store(key, concreteContainer{
		concrete: concrete,
		shared:   singleton,
		instance: nil,
		resolved: false,
	})

	// delete the alias if you explicitly bind
	delete(aliasList, key)

	// register DI provide if abstract is a non-string type, and it's singleton
	if singleton && key.isTypeKey {
		_ = registerProvider[VT](key)
	}

	// If the abstract type was already resolved in this container we'll fire the
	// rebound listener so that any objects which have already gotten resolved
	// can have their copy of the object updated via the listener callbacks.
	if _resolved {
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

	concrete, loaded := bindingList.LoadOrStore(key, concreteContainer{
		concrete: nil,
		shared:   true,
		instance: instance,
		resolved: true,
	})

	// key exists, update map
	if loaded {
		c := concrete.(concreteContainer) // assert will create a new struct
		c.instance = instance
		c.resolved = true
		c.shared = true

		// it cannot use the CompareAndSwap because concreteContainer.concrete is uncomparable type
		bindingList.Store(key, c)
	}

	return instance
}

func resolved(key tKey) bool {
	if concrete, ok := bindingList.Load(key); ok {
		return concrete.(concreteContainer).resolved
	}
	return false
}

func Resolved[VT VTConstraint](abstract VT) bool {
	key := toTKey[VT](abstract)
	key = getAlias(key)

	return resolved(key)
}

func IsShared[VT VTConstraint](abstract VT) bool {
	key := toTKey[VT](abstract)
	key = getAlias(key)

	if concrete, ok := bindingList.Load(key); ok {
		return concrete.(concreteContainer).shared
	}

	return false
}

// resolve the given type from the container.
func resolve[T any](key tKey, args ...any) (T, error) {
	var obj any
	var ok bool
	var err error
	var instance T

	concrete, loaded := bindingList.Load(key)

	if loaded {
		c := concrete.(concreteContainer)
		// singleton && resolved
		if c.shared && c.resolved {
			obj = c.instance
		} else { // first or non-singleton
			if c.concrete == nil {
				return instance, errors.Errorf("the concrete \"%s\" is invalid in the container", key.String())
			}

			obj, err = c.concrete(args...)
			if err != nil {
				return instance, err
			}

			// If the requested type is registered as a singleton we'll want to cache off
			// the instanceList in "memory" so we can return it later without creating an
			// entirely new instance of an object on each subsequent request for it.
			if c.shared {
				c.instance = obj
			}
			c.resolved = true

			// it cannot use the CompareAndSwap because concreteContainer.concrete is uncomparable type
			bindingList.Store(key, c)
		}
	}

	if instance, ok = obj.(T); ok {
		return instance, nil
	}

	return instance, errors.Errorf("[Container]: expected %s key to be type %T but got %T", key.String(), new(T), obj)
}

// Make Resolve the given type from the container.
//
//	Example: container.Make[*QueueManager]("queue.manager")
func Make[T any, VT VTConstraint](abstract VT, args ...any) (T, error) {
	key := toTKey[VT](abstract)
	key = getAlias(key)
	return resolve[T](key, args...)
}

// MustMake Resolve the given type from the container or panic.
func MustMake[T any, VT VTConstraint](abstract VT, args ...any) T {
	instance, err := Make[T](abstract, args...)
	if err != nil {
		panic(err)
	}

	return instance
}

// MakeAs Resolve the actualType from the container. Avoid generic writing style
//
//	Example: container.MakeAs("queue.manager", (*QueueManager)(nil))
func MakeAs[T any, VT VTConstraint](abstract VT, actualType T, args ...any) (T, error) {
	return Make[T](abstract, args...)
}

// MustMakeAs Resolve the given type from the container or panic. Avoid generic writing style
func MustMakeAs[T any, VT VTConstraint](abstract VT, actualType T, args ...any) T {
	return MustMake[T](abstract, args...)
}

func rebound(key tKey) {
	_, _ = resolve[any](key)
}

// Invoke a func via DI(IoC)
func Invoke(fn any) error {
	return di.Invoke(fn)
}
