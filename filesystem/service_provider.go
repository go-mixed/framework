package filesystem

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/filesystem"
	"gopkg.in/go-mixed/framework.v1/filesystem/disk"
)

type ServiceProvider struct {
}

func (sp *ServiceProvider) Register() {

	container.Singleton((*FilesystemManager)(nil), func(args ...any) (any, error) {
		m := NewFilesystemManager()
		m.Extend("local", func(driverName string, args ...any) (filesystem.IStorage, error) {
			return disk.NewLocal(driverName)
		}).Extend("s3", func(driverName string, args ...any) (filesystem.IStorage, error) {
			return disk.NewS3(context.Background(), driverName)
		})

		return m, nil
	})
	container.Alias("filesystem.manager", (*FilesystemManager)(nil))

	container.Singleton(filesystem.IStorage(nil), func(args ...any) (any, error) {
		return container.MustMakeAs("filesystem.manager", (*FilesystemManager)(nil)).DefaultDriver()
	})
	container.Alias("filesystem", filesystem.IStorage(nil))
	container.Alias("filesystem.disk", filesystem.IStorage(nil))
}

func (sp *ServiceProvider) Boot() {

}
