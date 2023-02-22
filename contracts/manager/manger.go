package manager

type Concrete[T any] func(name string) (T, error)

type IManager[T any] interface {
	DefaultDriverName() string
	HasDriver(name string) bool
	Driver(name string) (T, error)
	MustDriver(name string) T
	RemoveDriver(name string)
	DefaultDriver() (T, error)
	MustDefaultDriver() T
}
