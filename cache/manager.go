package cache

import (
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/support/manager"

	"gopkg.in/go-mixed/framework.v1/contracts/cache"
	"gopkg.in/go-mixed/framework.v1/facades"
)

type StoreManager struct {
	manager.Manager[cache.IStore]
}

func NewStoreManager() *StoreManager {
	m := &StoreManager{}
	m.Manager = manager.MakeManager[cache.IStore](m.DefaultDriverName)
	return m
}

func (m *StoreManager) DefaultDriverName() string {
	return facades.Config.GetString("cache.default")
}

func (m *StoreManager) makeDriver(storeName string) (cache.IStore, error) {
	driver := facades.Config.GetString("config."+storeName+".driver", "memory")
	driverContainerName := "cache.drivers." + driver

	if container.Bound(driverContainerName) {
		instance, err := container.Make[cache.IStore](driverContainerName, storeName)
		if err != nil {
			color.Redf("[Cache] Initialize %s driver of store %s error: %v\n", driver, storeName, err)
			return nil, errors.Errorf("[Cache] Initialize %s driver of store %s error: %v\n", driver, storeName, err)
		}

		return instance.(cache.IStore), nil
	}

	color.Redf("[Cache] %s driver of cache is not defined.\n", driver)
	return nil, errors.Errorf("[Cache] %s driver of cache is not defined.\n", driver)
}

func (m *StoreManager) Store(storeName string) (cache.IStore, error) {
	return m.Driver(storeName)
}

func (m *StoreManager) extendStores(manager *StoreManager) {
	for name := range facades.Config.GetMap("stores") {
		manager.Extend(name, manager.makeDriver)
	}
}
