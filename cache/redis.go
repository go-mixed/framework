package cache

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/manager"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	cachecontract "gopkg.in/go-mixed/framework.v1/contracts/cache"
)

type Redis struct {
	ctx       context.Context
	storeName string
	prefix    string
	redis     *redis.Client
}

func NewRedis(storeName string, ctx context.Context) (*Redis, error) {
	connection := config.GetString("cache.stores." + storeName + ".connection")
	if connection == "" {
		connection = "default"
	}

	host := config.GetString("database.redis." + connection + ".host")
	if host == "" {
		return nil, errors.New("redis host in invalid")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + config.GetString("database.redis."+connection+".port"),
		Password: config.GetString("database.redis." + connection + ".password"),
		DB:       config.GetInt("database.redis." + connection + ".database"),
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, errors.WithMessage(err, "init connection error")
	}

	return &Redis{
		ctx:       ctx,
		storeName: storeName,
		prefix:    prefix(),
		redis:     client,
	}, nil
}

func (r *Redis) Store(storeName string) cachecontract.IStore {
	return container.MustMake[manager.IManager[cachecontract.IStore]]("cache.manager").MustDriver(storeName)
}

func (r *Redis) WithContext(ctx context.Context) cachecontract.IStore {
	store, _ := NewRedis(r.storeName, ctx)

	return store
}

// Add Store an item in the cache if the key does not exist.
func (r *Redis) Add(key string, value any, seconds time.Duration) bool {
	val, err := r.redis.SetNX(r.ctx, r.prefix+key, value, seconds).Result()
	if err != nil {
		return false
	}

	return val
}

// Forever Store an item in the cache indefinitely.
func (r *Redis) Forever(key string, value any) bool {
	if err := r.Put(key, value, 0); err != nil {
		return false
	}

	return true
}

// Forget Remove an item from the cache.
func (r *Redis) Forget(key string) bool {
	_, err := r.redis.Del(r.ctx, r.prefix+key).Result()

	if err != nil {
		return false
	}

	return true
}

// Flush Remove all items from the cache.
func (r *Redis) Flush() bool {
	res, err := r.redis.FlushAll(r.ctx).Result()

	if err != nil || res != "OK" {
		return false
	}

	return true
}

// Get Retrieve an item from the cache by key.
func (r *Redis) Get(key string, def any) any {
	val, err := r.redis.Get(r.ctx, r.prefix+key).Result()
	if err != nil {
		switch s := def.(type) {
		case func() any:
			return s()
		default:
			return def
		}
	}

	return val
}

func (r *Redis) GetBool(key string, def bool) bool {
	res := r.Get(key, def)
	if val, ok := res.(string); ok {
		return val == "1"
	}

	return res.(bool)
}

func (r *Redis) GetInt(key string, def int) int {
	res := r.Get(key, def)
	if val, ok := res.(string); ok {
		i, err := strconv.Atoi(val)
		if err != nil {
			return def
		}

		return i
	}

	return res.(int)
}

func (r *Redis) GetString(key string, def string) string {
	return r.Get(key, def).(string)
}

// Has Check an item exists in the cache.
func (r *Redis) Has(key string) bool {
	value, err := r.redis.Exists(r.ctx, r.prefix+key).Result()

	if err != nil || value == 0 {
		return false
	}

	return true
}

// Pull Retrieve an item from the cache and delete it.
func (r *Redis) Pull(key string, def any) any {
	res := r.Get(key, def)
	r.Forget(key)

	return res
}

// Put Store an item in the cache for a given number of seconds.
func (r *Redis) Put(key string, value any, seconds time.Duration) error {
	err := r.redis.Set(r.ctx, r.prefix+key, value, seconds).Err()
	if err != nil {
		return err
	}

	return nil
}

// Remember Get an item from the cache, or execute the given Closure and store the result.
func (r *Redis) Remember(key string, seconds time.Duration, callback func() any) (any, error) {
	val := r.Get(key, nil)

	if val != nil {
		return val, nil
	}

	val = callback()

	if err := r.Put(key, val, seconds); err != nil {
		return nil, err
	}

	return val, nil
}

// RememberForever Get an item from the cache, or execute the given Closure and store the result forever.
func (r *Redis) RememberForever(key string, callback func() any) (any, error) {
	val := r.Get(key, nil)

	if val != nil {
		return val, nil
	}

	val = callback()

	if err := r.Put(key, val, 0); err != nil {
		return nil, err
	}

	return val, nil
}
