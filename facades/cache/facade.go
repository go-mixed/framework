package cache

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/cache"
	"time"
)

func getDefaultStore() cache.IStore {
	return container.MustMake[cache.IStore]("cache")
}

func Store(storeName string) cache.IStore {
	return getDefaultStore().Store(storeName)
}

func WithContext(ctx context.Context) cache.IStore {
	return getDefaultStore().WithContext(ctx)
}

// Get Retrieve an item from the cache by key.
func Get(key string, def any) any {
	return getDefaultStore().Get(key, def)
}

func GetBool(key string, def bool) bool {
	return getDefaultStore().GetBool(key, def)
}

func GetInt(key string, def int) int {
	return getDefaultStore().GetInt(key, def)
}

func GetString(key string, def string) string {
	return getDefaultStore().GetString(key, def)
}

// Has Check an item exists in the cache.
func Has(key string) bool {
	return getDefaultStore().Has(key)
}

// Put Store an item in the cache for a given number of seconds.
func Put(key string, value any, sec time.Duration) error {
	return getDefaultStore().Put(key, value, sec)
}

// Pull Retrieve an item from the cache and delete it.
func Pull(key string, def any) any {
	return getDefaultStore().Pull(key, def)
}

// Add Store an item in the cache if the key does not exist.
func Add(key string, value any, sec time.Duration) bool {
	return getDefaultStore().Add(key, value, sec)
}

// Remember Get an item from the cache, or execute the given Closure and store the result.
func Remember(key string, ttl time.Duration, callback func() any) (any, error) {
	return getDefaultStore().Remember(key, ttl, callback)
}

// RememberForever Get an item from the cache, or execute the given Closure and store the result forever.
func RememberForever(key string, callback func() any) (any, error) {
	return getDefaultStore().RememberForever(key, callback)
}

// Forever Store an item in the cache indefinitely.
func Forever(key string, value any) bool {
	return getDefaultStore().Forever(key, value)
}

// Forget Remove an item from the cache.
func Forget(key string) bool {
	return getDefaultStore().Forget(key)
}

// Flush Remove all items from the cache.
func Flush() bool {
	return getDefaultStore().Flush()
}
