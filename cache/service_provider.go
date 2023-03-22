package cache

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/cache/console"
	"gopkg.in/go-mixed/framework.v1/cache/store"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts"
	"gopkg.in/go-mixed/framework.v1/contracts/cache"
	console2 "gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
)

type ServiceProvider struct {
}

var _ contracts.IServiceProvider = (*ServiceProvider)(nil)

func (sp *ServiceProvider) Register() {
	container.Singleton((*CacheManager)(nil), func(args ...any) (any, error) {
		m := NewCacheManager()
		m.Extend("redis", func(driverName string, args ...any) (cache.IStore, error) {
			return store.NewRedis(driverName, context.Background())
		}).Extend("memory", func(driverName string, args ...any) (cache.IStore, error) {
			return store.NewMemory()
		})

		return m, nil
	})
	container.Alias("cache.manager", (*CacheManager)(nil))

	container.Singleton(cache.IStore(nil), func(args ...any) (any, error) {
		return container.MustMakeAs("cache.manager", (*CacheManager)(nil)).DefaultDriver()
	})
	container.Alias("cache", cache.IStore(nil))
	container.Alias("cache.store", cache.IStore(nil))
}

func (sp *ServiceProvider) Boot() {
	sp.registerCommands()
}

func (sp *ServiceProvider) registerCommands() {
	artisan.Register([]console2.ICommand{
		&console.ClearCommand{},
	})
}
