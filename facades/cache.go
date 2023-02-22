package facades

import (
	"context"
	"gopkg.in/go-mixed/framework/cache"
	cachecontract "gopkg.in/go-mixed/framework/contracts/cache"
)

var Cache *cachecontract.IStore = cache.Get

func (c *cacheFacade) Store(storeName string) cachecontract.IStore {
	return GetCacheManager().Store(storeName)
}

func (c *cacheFacade) WithContext(ctx context.Context) cachecontract.IStore {
	return GetDefaultStore().WithContext(ctx)
}

func (c *cacheFacade) Get(key string, def any) any {
	return GetDefaultStore().Get(key, def)
}
