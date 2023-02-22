package container

import (
	"errors"
	"gopkg.in/go-mixed/framework/contracts/container"
)

var ErrEntryNotExists = errors.New("abstract is not exists in the application")

type invoker struct {
	// true for singleton or instances
	shared   bool
	concrete container.Concrete
}

type Container struct {
	// The container's bindings.
	bindings  map[string]invoker
	instances map[string]any
	resolved  map[string]bool
}

var _ container.IContainer = (*Container)(nil)

func Initial() Container {
	return Container{
		bindings:  map[string]invoker{},
		instances: map[string]any{},
		resolved:  map[string]bool{},
	}
}

// Bound Determine if the given abstract type has been bound.
func (c *Container) Bound(abstract string) bool {
	if _, ok := c.bindings[abstract]; ok {
		return ok
	}

	if _, ok := c.instances[abstract]; ok {
		return ok
	}

	return false
}

func (c *Container) Has(abstract string) bool {
	return c.Bound(abstract)
}

func (c *Container) Bind(abstract string, concrete container.Concrete, shared bool) container.IContainer {
	c.bindings[abstract] = invoker{
		shared,
		concrete,
	}

	// If the abstract type was already resolved in this container we'll fire the
	// rebound listener so that any objects which have already gotten resolved
	// can have their copy of the object updated via the listener callbacks.
	if c.Resolved(abstract) {
		c.Rebound(abstract)
	}

	return c
}

func (c *Container) Singleton(abstract string, concrete container.Concrete) container.IContainer {
	return c.Bind(abstract, concrete, true)
}

// Instance Register an existing instance as shared in the container.
func (c *Container) Instance(abstract string, instance any) any {

	if fn, ok := instance.(func() any); ok {
		return c.Bind(abstract, func(args ...any) (any, error) {
			return fn(), nil
		}, false)
	}

	// We'll check to determine if this type has been bound before, and if it has
	// we will fire the rebound callbacks registered with the container, and it
	// can be updated with consuming classes that have gotten resolved here.
	c.instances[abstract] = instance

	return instance
}

func (c *Container) Resolved(abstract string) bool {
	if _, ok := c.resolved[abstract]; ok {
		return true
	}
	_, ok := c.instances[abstract]
	return ok
}

func (c *Container) IsShared(abstract string) bool {
	if _, ok := c.instances[abstract]; ok {
		return ok
	}

	i, ok := c.bindings[abstract]
	return ok && i.shared
}

// Resolve the given type from the container.
func (c *Container) Resolve(abstract string, args ...any) (any, error) {
	if obj, ok := c.instances[abstract]; ok {
		return obj, nil
	}

	_invoker, ok := c.bindings[abstract]
	if !ok {
		return nil, ErrEntryNotExists
	}

	instance, err := _invoker.concrete(args...)
	if err != nil {
		return nil, err
	}

	// If the requested type is registered as a singleton we'll want to cache off
	// the instances in "memory" so we can return it later without creating an
	// entirely new instance of an object on each subsequent request for it.
	if _invoker.shared {
		c.instances[abstract] = instance
	}

	c.resolved[abstract] = true
	return instance, nil
}

// Make Resolve the given type from the container.
func (c *Container) Make(abstract string, args ...any) (any, error) {
	return c.Resolve(abstract, args...)
}

func (c *Container) MustMake(abstract string, args ...any) any {
	instance, err := c.Make(abstract, args...)
	if err != nil {
		panic(err)
	}

	return instance
}

func (c *Container) MakeT(abstract string, args ...any) container.InstanceResult {
	instance, err := c.Make(abstract, args...)
	return container.InstanceResult{Instance: instance, Error: err}
}

func (c *Container) Rebound(abstract string) {
	_, _ = c.Make(abstract)
}
