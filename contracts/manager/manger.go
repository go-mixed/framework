package manager

type Concrete[T any] func(driverName string, args ...any) (T, error)

type IManager[T any] interface {
	DefaultDriverName() string
	HasCustomCreator(name string) bool
	Resolved(name string) bool
	Driver(name string) (T, error)
	MustDriver(name string) T
	Remove(name string)
	RemoveCustomCreator(name string)
	DefaultDriver() (T, error)
	MustDefaultDriver() T
}
