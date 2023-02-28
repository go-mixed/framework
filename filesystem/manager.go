package filesystem

import (
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"gopkg.in/go-mixed/framework.v1/contracts/filesystem"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"gopkg.in/go-mixed/framework.v1/support/manager"
)

type FilesystemManager struct {
	manager.Manager[filesystem.IStorage]
}

func NewFilesystemManager() *FilesystemManager {
	m := &FilesystemManager{}
	m.Manager = manager.MakeManager[filesystem.IStorage](m.DefaultDiskName, m.makeDisk)
	return m
}

func (m *FilesystemManager) DefaultDiskName() string {
	return config.GetString("filesystems.default")
}

func (m *FilesystemManager) Disk(name string) filesystem.IStorage {
	return m.Manager.MustDriver(name)
}

func (m *FilesystemManager) makeDisk(diskName string) (filesystem.IStorage, error) {
	driver := config.GetString("filesystems.disks."+diskName+".driver", "")

	if m.HasCustomCreator(driver) {
		instance, err := m.CallCustomCreator(diskName, driver)
		if err != nil {
			color.Redf("[Cache] Initialize filesystem driver \"%s.%s\" error: %v\n", diskName, driver, err)
			return nil, errors.Errorf("[Cache] Initialize filesystem driver \"%s.%s\" error: %v\n", diskName, driver, err)
		}

		return instance.(filesystem.IStorage), nil
	}

	color.Redf("[Cache] filesystem driver \"%s.%s\" is not defined.\n", diskName, driver)
	return nil, errors.Errorf("[Cache] filesystem driver \"%s.%s\" is not defined.\n", diskName, driver)
}
