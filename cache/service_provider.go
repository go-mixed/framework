package cache

import (
	"context"
	"gopkg.in/go-mixed/framework/cache/console"
	"gopkg.in/go-mixed/framework/contracts"
	console2 "gopkg.in/go-mixed/framework/contracts/console"
	"gopkg.in/go-mixed/framework/contracts/container"
	"gopkg.in/go-mixed/framework/facades"
)

type ServiceProvider struct {
}

var _ contracts.IServiceProvider = (*ServiceProvider)(nil)

func (sp *ServiceProvider) Register(container container.IContainer) {
	container.Instance("cache.manager", func() any {
		manager := NewStoreManager(container)
		sp.registerStores(manager)
		return manager
	})

	sp.registerCacheDrivers(container)

}

func (sp *ServiceProvider) registerStores(manager *StoreManager) {
	for name := range facades.Config.GetMap("stores") {
		manager.Extend(name, manager.makeDriver)
	}
}

func (sp *ServiceProvider) registerCacheDrivers(container container.IContainer) {
	container.Bind("cache.drivers.redis", func(args ...any) (any, error) {
		return NewRedis(args[0].(string), context.Background())
	}, false)

	container.Bind("cache.drivers.memory", func(args ...any) (any, error) {
		return NewMemory()
	}, false)

}

func (sp *ServiceProvider) Boot(container container.IContainer) {
	sp.registerCommands()
}

func (sp *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]console2.Command{
		&console.ClearCommand{},
	})
}
