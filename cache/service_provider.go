package cache

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/cache/console"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts"
	"gopkg.in/go-mixed/framework.v1/contracts/cache"
	console2 "gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/facades"
)

type ServiceProvider struct {
}

var _ contracts.IServiceProvider = (*ServiceProvider)(nil)

func (sp *ServiceProvider) Register() {
	container.Singleton((*StoreManager)(nil), func(args ...any) (any, error) {
		manager := NewStoreManager()
		manager.extendStores(manager)
		return manager, nil
	})
	container.Alias("cache.manager", (*StoreManager)(nil))

	container.Singleton(cache.IStore(nil), func(args ...any) (any, error) {
		return container.MustMake[*StoreManager]("cache.manager").DefaultDriver()
	})
	container.Alias("cache.store", cache.IStore(nil))

	sp.registerCacheDrivers()
}

func (sp *ServiceProvider) registerCacheDrivers() {
	container.Bind("cache.drivers.redis", func(args ...any) (any, error) {
		return NewRedis(args[0].(string), context.Background())
	}, false)

	container.Bind("cache.drivers.memory", func(args ...any) (any, error) {
		return NewMemory()
	}, false)

}

func (sp *ServiceProvider) Boot() {
	sp.registerCommands()
}

func (sp *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]console2.Command{
		&console.ClearCommand{},
	})
}
