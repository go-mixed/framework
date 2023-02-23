package filesystem

import (
	"context"
	"fmt"
	"gopkg.in/go-mixed/framework.v1/facades/config"

	"gopkg.in/go-mixed/framework.v1/contracts/filesystem"
	"gopkg.in/go-mixed/framework.v1/facades"
)

type Driver string

const (
	DriverLocal  Driver = "local"
	DriverS3     Driver = "s3"
	DriverCustom Driver = "custom"
)

type Storage struct {
	filesystem.Driver
	drivers map[string]filesystem.Driver
}

func NewStorage() *Storage {
	defaultDisk := config.GetString("filesystems.default")
	if defaultDisk == "" {
		facades.Log.Errorf("[filesystem] please set default disk")
		return nil
	}

	driver, err := NewDriver(defaultDisk)
	if err != nil {
		facades.Log.Errorf("[filesystem] init %s disk error: %+v", defaultDisk, err)
		return nil
	}

	drivers := make(map[string]filesystem.Driver)
	drivers[defaultDisk] = driver
	return &Storage{
		Driver:  driver,
		drivers: drivers,
	}
}

func NewDriver(disk string) (filesystem.Driver, error) {
	ctx := context.Background()
	driver := Driver(config.GetString(fmt.Sprintf("filesystems.disks.%s.driver", disk)))
	switch driver {
	case DriverLocal:
		return NewLocal(disk)
	case DriverS3:
		return NewS3(ctx, disk)
	case DriverCustom:
		driver, ok := config.Get(fmt.Sprintf("filesystems.disks.%s.via", disk)).(filesystem.Driver)
		if !ok {
			return nil, fmt.Errorf("[filesystem] init %s disk fail: via must be filesystem.Driver.", disk)
		}

		return driver, nil
	}

	return nil, fmt.Errorf("[filesystem] invalid driver: %s, only support local, s3, oss, cos, custom.", driver)
}

func (r *Storage) Disk(disk string) filesystem.Driver {
	if driver, exist := r.drivers[disk]; exist {
		return driver
	}

	driver, err := NewDriver(disk)
	if err != nil {
		facades.Log.Error(err.Error())

		return nil
	}

	r.drivers[disk] = driver

	return driver
}
