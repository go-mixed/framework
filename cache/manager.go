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
	m.Manager = manager.MakeManager[cache.IStore](m.DefaultCacheName, m.makeStore)
	return m
}

func (m *CacheManager) DefaultCacheName() string {
	return config.GetString("cache.default")
}

func (m *CacheManager) makeStore(storeName string) (cache.IStore, error) {
	driver := config.GetString("cache.stores."+storeName+".driver", "")

	if m.HasCustomCreator(driver) {
		instance, err := m.CallCustomCreator(driver, storeName)
		if err != nil {
			color.Redf("[Cache] Initialize cache driver \"%s.%s\" error: %v\n", storeName, driver, err)
			return nil, errors.Errorf("[Cache] Initialize cache \"%s.%s\" error: %v\n", storeName, driver, err)
		}

		return instance.(cache.IStore), nil
	}

	color.Redf("[Cache] cache driver \"%s.%s\" is not defined.\n", storeName, driver)
	return nil, errors.Errorf("[Cache] cache driver \"%s.%s\" is not defined.\n", storeName, driver)
}

func (m *CacheManager) Store(storeName string) (cache.IStore, error) {
	return m.Driver(storeName)
}
