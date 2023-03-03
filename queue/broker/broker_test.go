package broker

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configmocks "gopkg.in/go-mixed/framework.v1/contracts/config/mocks"
	"gopkg.in/go-mixed/framework.v1/contracts/event"
	"gopkg.in/go-mixed/framework.v1/testing/mock"
)

func TestSync(t *testing.T) {
	mockConfig := mock.Config()
	mockConfig.On("GetString", "queue.connections.sync.driver").Return("sync").Once()

	server := NewSyncBroker("sync")
	assert.Nil(t, server)

	mockConfig.AssertExpectations(t)
}

func TestRedis(t *testing.T) {
	mockConfig := mock.Config()
	mockConfig.On("GetString", "queue.connections.redis.driver").Return("redis").Once()
	mockConfig.On("GetString", "queue.connections.redis.connection").Return("default").Once()
	mockConfig.On("GetString", "database.redis.default.host").Return("127.0.0.1").Once()
	mockConfig.On("GetString", "database.redis.default.password").Return("").Once()
	mockConfig.On("GetString", "database.redis.default.port").Return("6379").Once()
	mockConfig.On("GetInt", "database.redis.default.database").Return(0).Once()
	mockConfig.On("GetString", "queue.connections.redis.queue", "default").Return("default").Once()
	mockConfig.On("GetString", "app.name").Return("goravel").Once()

	server, _ := NewRedisBroker("redis")
	assert.NotNil(t, server)

	mockConfig.AssertExpectations(t)
}

func TestGetQueueName(t *testing.T) {
	var (
		mockConfig *configmocks.Config
	)

	beforeEach := func() {
		mockConfig = mock.Config()
	}

	tests := []struct {
		description     string
		setup           func()
		connection      string
		queue           string
		expectQueueName string
	}{
		{
			description: "success when connection and queue are empty",
			setup: func() {
				mockConfig.On("GetString", "app.name").Return("").Once()
				mockConfig.On("GetString", "queue.default").Return("redis").Once()
				mockConfig.On("GetString", "queue.connections.redis.queue", "default").Return("queue").Once()
			},
			expectQueueName: "laravel_queues:queue",
		},
		{
			description: "success when connection and queue aren't empty",
			setup: func() {
				mockConfig.On("GetString", "app.name").Return("app").Once()

			},
			connection:      "redis",
			queue:           "queue",
			expectQueueName: "app_queues:queue",
		},
	}

	for _, test := range tests {
		beforeEach()
		test.setup()
		queueName := GetQueueName(test.connection, test.queue)
		assert.Equal(t, test.expectQueueName, queueName, test.description)
	}
}

func TestGetRedisConfig(t *testing.T) {
	mockConfig := mock.Config()
	mockConfig.On("GetString", "queue.connections.redis.connection").Return("default").Once()
	mockConfig.On("GetString", "database.redis.default.host").Return("127.0.0.1").Once()
	mockConfig.On("GetString", "database.redis.default.password").Return("").Once()
	mockConfig.On("GetString", "database.redis.default.port").Return("6379").Once()
	mockConfig.On("GetInt", "database.redis.default.database").Return(0).Once()
	mockConfig.On("GetString", "queue.connections.redis.queue", "default").Return("default").Once()
	mockConfig.On("GetString", "app.name").Return("goravel").Once()

	redisConfig, database, queue := getRedisConfig("redis")
	assert.Equal(t, "127.0.0.1:6379", redisConfig)
	assert.Equal(t, 0, database)
	assert.Equal(t, "laravel_queues:default", queue)
}

type TestJob struct {
}

func (receiver *TestJob) Signature() string {
	return "TestName"
}

func (receiver *TestJob) Handle(args ...any) error {
	return nil
}

type TestJobDuplicate struct {
}

func (receiver *TestJobDuplicate) Signature() string {
	return "TestName"
}

func (receiver *TestJobDuplicate) Handle(args ...any) error {
	return nil
}

type TestJobEmpty struct {
}

func (receiver *TestJobEmpty) Signature() string {
	return ""
}

func (receiver *TestJobEmpty) Handle(args ...any) error {
	return nil
}

type TestEvent struct {
}

func (receiver *TestEvent) Signature() string {
	return "TestName"
}

func (receiver *TestEvent) Handle(args []event.Arg) ([]event.Arg, error) {
	return args, nil
}

type TestListener struct {
}

func (receiver *TestListener) Signature() string {
	return "TestName"
}

func (receiver *TestListener) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     false,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *TestListener) Handle(args ...any) error {
	return nil
}

type TestListenerDuplicate struct {
}

func (receiver *TestListenerDuplicate) Signature() string {
	return "TestName"
}

func (receiver *TestListenerDuplicate) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     false,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *TestListenerDuplicate) Handle(args ...any) error {
	return nil
}

type TestListenerEmpty struct {
}

func (receiver *TestListenerEmpty) Signature() string {
	return ""
}

func (receiver *TestListenerEmpty) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     false,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *TestListenerEmpty) Handle(args ...any) error {
	return nil
}

func TestEvents2Tasks(t *testing.T) {
	_, err := eventsToJobMap(map[event.Event][]event.Listener{
		&TestEvent{}: {
			&TestListener{},
		},
	})
	assert.Nil(t, err)

	_, err = eventsToJobMap(map[event.Event][]event.Listener{
		&TestEvent{}: {
			&TestListener{},
			&TestListenerDuplicate{},
		},
	})
	assert.Nil(t, err)

	_, err = eventsToJobMap(map[event.Event][]event.Listener{
		&TestEvent{}: {
			&TestListenerEmpty{},
		},
	})

	assert.NotNil(t, err)
}
