package cache

import (
	cachecontract "gopkg.in/go-mixed/framework.v1/contracts/cache"
	"gopkg.in/go-mixed/framework.v1/support"
)

func GetCacheManager() *StoreManager {
	manager, err := support.MustAs("cache.manager", nil, (*StoreManager)(nil))
	if err != nil {
		panic(err)
	}

	return manager
}

func GetDefaultStore() cachecontract.IStore {
	store, err := support.MustMakeAs("cache.manager", nil, cachecontract.IStore(nil))
	if err != nil {
		panic(err)
	}

	return store
}
