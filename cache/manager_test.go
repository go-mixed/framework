package cache

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/cache/store"
	"log"
	"testing"
	"time"

	"gopkg.in/go-mixed/framework.v1/contracts/cache"
	testingdocker "gopkg.in/go-mixed/framework.v1/testing/docker"
	"gopkg.in/go-mixed/framework.v1/testing/mock"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/suite"
)

type ModuleTestSuite struct {
	suite.Suite
	stores      map[string]cache.IStore
	redisDocker *dockertest.Resource
}

func TestApplicationTestSuite(t *testing.T) {
	redisPool, redisDocker, redisStore, err := getRedisDocker()
	if err != nil {
		log.Fatalf("Get redis store error: %s", err)
	}
	memoryStore, err := getMemoryStore()
	if err != nil {
		log.Fatalf("Get memory store error: %s", err)
	}

	suite.Run(t, &ModuleTestSuite{
		stores: map[string]cache.IStore{
			"redis":  redisStore,
			"memory": memoryStore,
		},
		redisDocker: redisDocker,
	})

	if err := redisPool.Purge(redisDocker); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func (s *ModuleTestSuite) SetupTest() {
}

func (s *ModuleTestSuite) TestInitRedis() {
	tests := []struct {
		description string
		setup       func(description string)
	}{
		{
			description: "success",
			setup: func(description string) {
				mockConfig := mock.Config()
				mockConfig.On("GetString", "cache.default").Return("redis").Twice()
				mockConfig.On("GetString", "cache.stores.redis.driver").Return("redis").Once()
				mockConfig.On("GetString", "cache.stores.redis.connection").Return("default").Once()
				mockConfig.On("GetString", "database.redis.default.host").Return("localhost").Once()
				mockConfig.On("GetString", "database.redis.default.port").Return(s.redisDocker.GetPort("6379/tcp")).Once()
				mockConfig.On("GetString", "database.redis.default.password").Return("").Once()
				mockConfig.On("GetInt", "database.redis.default.database").Return(0).Once()
				mockConfig.On("GetString", "cache.prefix").Return("laravel_cache").Once()

				manager := CacheManager{}
				s.NotNil(manager, description)

				mockConfig.AssertExpectations(s.T())
			},
		},
		{
			description: "error",
			setup: func(description string) {
				mockConfig := mock.Config()
				mockConfig.On("GetString", "cache.default").Return("redis").Twice()
				mockConfig.On("GetString", "cache.stores.redis.driver").Return("redis").Once()
				mockConfig.On("GetString", "cache.stores.redis.connection").Return("default").Once()
				mockConfig.On("GetString", "database.redis.default.host").Return("").Once()

				manager := CacheManager{}
				s.Nil(manager, description)

				mockConfig.AssertExpectations(s.T())
			},
		},
	}

	for _, test := range tests {
		test.setup(test.description)
	}
}

func (s *ModuleTestSuite) TestAdd() {
	for name, store := range s.stores {
		s.Run(name, func() {
			s.Nil(store.Put("name", "Laravel", 1*time.Second))
			s.False(store.Add("name", "World", 1*time.Second))
			s.True(store.Add("name1", "World", 1*time.Second))
			s.True(store.Has("name1"))
			time.Sleep(2 * time.Second)
			s.False(store.Has("name1"))
			s.True(store.Flush())
		})
	}
}

func (s *ModuleTestSuite) TestForever() {
	for name, store := range s.stores {
		s.Run(name, func() {
			s.True(store.Forever("name", "Laravel"))
			s.Equal("Laravel", store.Get("name", "").(string))
			s.True(store.Flush())
		})
	}
}

func (s *ModuleTestSuite) TestForget() {
	for name, store := range s.stores {
		s.Run(name, func() {
			val := store.Forget("test-forget")
			s.True(val)

			err := store.Put("test-forget", "laravel", 5*time.Second)
			s.Nil(err)
			s.True(store.Forget("test-forget"))
		})
	}
}

func (s *ModuleTestSuite) TestFlush() {
	for name, store := range s.stores {
		s.Run(name, func() {
			s.Nil(store.Put("test-flush", "laravel", 5*time.Second))
			s.Equal("laravel", store.Get("test-flush", nil).(string))

			s.True(store.Flush())
			s.False(store.Has("test-flush"))
		})
	}
}

func (s *ModuleTestSuite) TestGet() {
	for name, store := range s.stores {
		s.Run(name, func() {
			s.Nil(store.Put("name", "Laravel", 1*time.Second))
			s.Equal("Laravel", store.Get("name", "").(string))
			s.Equal("World", store.Get("name1", "World").(string))
			s.Equal("World1", store.Get("name2", func() any {
				return "World1"
			}).(string))
			s.True(store.Forget("name"))
			s.True(store.Flush())
		})
	}
}

func (s *ModuleTestSuite) TestGetBool() {
	for name, store := range s.stores {
		s.Run(name, func() {
			s.Equal(true, store.GetBool("test-get-bool", true))
			s.Nil(store.Put("test-get-bool", true, 2*time.Second))
			s.Equal(true, store.GetBool("test-get-bool", false))
		})
	}
}

func (s *ModuleTestSuite) TestGetInt() {
	for name, store := range s.stores {
		s.Run(name, func() {
			s.Equal(2, store.GetInt("test-get-int", 2))
			s.Nil(store.Put("test-get-int", 3, 2*time.Second))
			s.Equal(3, store.GetInt("test-get-int", 2))
		})
	}
}

func (s *ModuleTestSuite) TestGetString() {
	for name, store := range s.stores {
		s.Run(name, func() {
			s.Equal("2", store.GetString("test-get-string", "2"))
			s.Nil(store.Put("test-get-string", "3", 2*time.Second))
			s.Equal("3", store.GetString("test-get-string", "2"))
		})
	}
}

func (s *ModuleTestSuite) TestHas() {
	for name, store := range s.stores {
		s.Run(name, func() {
			s.False(store.Has("test-has"))
			s.Nil(store.Put("test-has", "laravel", 5*time.Second))
			s.True(store.Has("test-has"))
		})
	}
}

func (s *ModuleTestSuite) TestPull() {
	for name, store := range s.stores {
		s.Run(name, func() {
			s.Nil(store.Put("name", "Laravel", 1*time.Second))
			s.True(store.Has("name"))
			s.Equal("Laravel", store.Pull("name", "").(string))
			s.False(store.Has("name"))
		})
	}
}

func (s *ModuleTestSuite) TestPut() {
	for name, store := range s.stores {
		s.Run(name, func() {
			s.Nil(store.Put("name", "Laravel", 1*time.Second))
			s.True(store.Has("name"))
			s.Equal("Laravel", store.Get("name", "").(string))
			time.Sleep(2 * time.Second)
			s.False(store.Has("name"))
		})
	}
}

func (s *ModuleTestSuite) TestRemember() {
	for name, store := range s.stores {
		s.Run(name, func() {
			s.Nil(store.Put("name", "Laravel", 1*time.Second))
			value, err := store.Remember("name", 1*time.Second, func() any {
				return "World"
			})
			s.Nil(err)
			s.Equal("Laravel", value)

			value, err = store.Remember("name1", 1*time.Second, func() any {
				return "World1"
			})
			s.Nil(err)
			s.Equal("World1", value)
			time.Sleep(2 * time.Second)
			s.False(store.Has("name1"))
			s.True(store.Flush())
		})
	}
}

func (s *ModuleTestSuite) TestRememberForever() {
	for name, store := range s.stores {
		s.Run(name, func() {
			s.Nil(store.Put("name", "Laravel", 1*time.Second))
			value, err := store.RememberForever("name", func() any {
				return "World"
			})
			s.Nil(err)
			s.Equal("Laravel", value)

			value, err = store.RememberForever("name1", func() any {
				return "World1"
			})
			s.Nil(err)
			s.Equal("World1", value)
			s.True(store.Flush())
		})
	}
}

func (s *ModuleTestSuite) TestCustomDriver() {
	mockConfig := mock.Config()
	mockConfig.On("GetString", "cache.default").Return("store").Once()
	mockConfig.On("GetString", "cache.stores.store.driver").Return("custom").Once()
	mockConfig.On("Get", "cache.stores.store.via").Return(&Store{}).Once()

	manager := CacheManager{}
	store := manager.MustDefaultDriver()
	s.NotNil(store)
	s.Equal("Laravel", store.Get("name", "Laravel").(string))

	mockConfig.AssertExpectations(s.T())
}

func getRedisDocker() (*dockertest.Pool, *dockertest.Resource, cache.IStore, error) {
	pool, resource, err := testingdocker.Redis()
	if err != nil {
		return nil, nil, nil, err
	}

	var istore cache.IStore
	if err := pool.Retry(func() error {
		var err error
		mockConfig := mock.Config()
		mockConfig.On("GetString", "cache.default").Return("redis").Once()
		mockConfig.On("GetString", "cache.stores.redis.connection").Return("default").Once()
		mockConfig.On("GetString", "database.redis.default.host").Return("localhost").Once()
		mockConfig.On("GetString", "database.redis.default.port").Return(resource.GetPort("6379/tcp")).Once()
		mockConfig.On("GetString", "database.redis.default.password").Return(resource.GetPort("")).Once()
		mockConfig.On("GetInt", "database.redis.default.database").Return(0).Once()
		mockConfig.On("GetString", "cache.prefix").Return("laravel_cache").Once()
		istore, err = store.NewRedis("redis", context.Background())

		return err
	}); err != nil {
		return nil, nil, nil, err
	}

	return pool, resource, istore, nil
}

func getMemoryStore() (*store.Memory, error) {
	mockConfig := mock.Config()
	mockConfig.On("GetString", "cache.prefix").Return("laravel_cache").Once()

	memory, err := store.NewMemory()
	if err != nil {
		return nil, err
	}

	return memory, nil
}

type Store struct {
}

func (r *Store) Store(storeName string) cache.IStore {
	return r
}

func (r *Store) WithContext(ctx context.Context) cache.IStore {
	return r
}

// Get Retrieve an item from the cache by key.
func (r *Store) Get(key string, def any) any {
	return def
}

// Get Retrieve an item from the cache by key.
func (r *Store) GetInt(key string, def int) int {
	return def
}

// Get Retrieve an item from the cache by key.
func (r *Store) GetBool(key string, def bool) bool {
	return def
}

// Get Retrieve an item from the cache by key.
func (r *Store) GetString(key string, def string) string {
	return def
}

// Has Check an item exists in the cache.
func (r *Store) Has(key string) bool {
	return true
}

// Put Store an item in the cache for a given number of seconds.
func (r *Store) Put(key string, value any, seconds time.Duration) error {
	return nil
}

// Pull Retrieve an item from the cache and delete it.
func (r *Store) Pull(key string, def any) any {
	return def
}

// Add Store an item in the cache if the key does not exist.
func (r *Store) Add(key string, value any, seconds time.Duration) bool {
	return true
}

// Remember Get an item from the cache, or execute the given Closure and store the result.
func (r *Store) Remember(key string, ttl time.Duration, callback func() any) (any, error) {
	return "", nil
}

// RememberForever Get an item from the cache, or execute the given Closure and store the result forever.
func (r *Store) RememberForever(key string, callback func() any) (any, error) {
	return "", nil
}

// Forever Store an item in the cache indefinitely.
func (r *Store) Forever(key string, value any) bool {
	return true
}

// Forget Remove an item from the cache.
func (r *Store) Forget(key string) bool {
	return true
}

// Flush Remove all items from the cache.
func (r *Store) Flush() bool {
	return true
}

var _ cache.IStore = &Store{}
