package cache

import (
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"gopkg.in/go-mixed/framework/contracts/container"
	manager2 "gopkg.in/go-mixed/framework/contracts/manager"
	"gopkg.in/go-mixed/framework/support"
	"gopkg.in/go-mixed/framework/support/manager"

	"gopkg.in/go-mixed/framework/contracts/cache"
	"gopkg.in/go-mixed/framework/facades"
)

type StoreManager struct {
	container container.IContainer

	manager.Manager[cache.IStore]
}

var _ manager2.IManager[cache.IStore] = (*StoreManager)(nil)

func NewStoreManager(container container.IContainer) *StoreManager {
	m := &StoreManager{container: container}
	m.Manager = manager.MakeManager[cache.IStore](m.DefaultDriverName)
	return m
}

func (m *StoreManager) DefaultDriverName() string {
	return facades.Config.GetString("cache.default")
}

func (m *StoreManager) makeDriver(storeName string) (cache.IStore, error) {
	driver := facades.Config.GetString("config."+storeName+".driver", "memory")
	driverContainerName := "cache.drivers." + driver
	if m.container.Bound(driverContainerName) {
		instance, err := support.As(m.container.MakeT(driverContainerName, storeName), cache.IStore(nil))
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
