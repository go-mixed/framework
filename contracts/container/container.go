package container

type Concrete func(args ...any) (any, error)

type InstanceResult struct {
	Instance any
	Error    error
}

type IContainer interface {
	Bound(abstract string) bool
	Has(abstract string) bool
	Bind(abstract string, concrete Concrete, shared bool) IContainer
	Singleton(abstract string, concrete Concrete) IContainer
	Instance(abstract string, instance any) any
	Resolved(abstract string) bool
	IsShared(abstract string) bool
	Resolve(abstract string, args ...any) (any, error)
	Make(abstract string, args ...any) (any, error)
	MustMake(abstract string, args ...any) any
	Rebound(abstract string)
}
