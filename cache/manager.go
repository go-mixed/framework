package cache

import (
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"gopkg.in/go-mixed/framework.v1/support/manager"

	"gopkg.in/go-mixed/framework.v1/contracts/cache"
)

type CacheManager struct {
	manager.Manager[cache.IStore]
}

func NewCacheManager() *CacheManager {
	m := &CacheManager{}
	m.Manager = manager.MakeManager[cache.IStore](m.DefaultDriverName, m.makeStore)
	return m
}

func (m *CacheManager) DefaultDriverName() string {
	return config.GetString("cache.default")
}

func (m *CacheManager) makeStore(storeName string) (cache.IStore, error) {
	driver := config.GetString("config.stores."+storeName+".driver", "memory")

	if m.HasCustomCreator(driver) {
		instance, err := m.CallCustomCreator(driver, storeName)
		if err != nil {
			color.Redf("[Cache] Initialize %s driver of store %s error: %v\n", driver, storeName, err)
			return nil, errors.Errorf("[Cache] Initialize %s driver of store %s error: %v\n", driver, storeName, err)
		}

		return instance.(cache.IStore), nil
	}

	color.Redf("[Cache] %s driver of cache is not defined.\n", driver)
	return nil, errors.Errorf("[Cache] %s driver of cache is not defined.\n", driver)
}

func (m *CacheManager) Store(storeName string) (cache.IStore, error) {
	return m.Driver(storeName)
}
