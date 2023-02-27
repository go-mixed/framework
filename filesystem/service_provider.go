package filesystem

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/filesystem"
)

type ServiceProvider struct {
}

func (sp *ServiceProvider) Register() {

	container.Singleton((*FilesystemManager)(nil), func(args ...any) (any, error) {
		m := NewFilesystemManager()
		m.Extend("local", func(driverName string, args ...any) (filesystem.IStorage, error) {
			return NewLocal(driverName)
		}).Extend("s3", func(driverName string, args ...any) (filesystem.IStorage, error) {
			return NewS3(context.Background(), driverName)
		})

		return m, nil
	})
	container.Alias("filesystem.manager", (*FilesystemManager)(nil))

	container.Singleton(filesystem.IStorage(nil), func(args ...any) (any, error) {
		return container.MustMake[*FilesystemManager]("filesystem.manager").DefaultDriver()
	})
	container.Alias("filesystem", filesystem.IStorage(nil))
	container.Alias("filesystem.disk", filesystem.IStorage(nil))
}

func (sp *ServiceProvider) Boot() {

}
