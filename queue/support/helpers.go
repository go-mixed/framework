package support

import (
	"errors"
	"fmt"
	"gopkg.in/go-mixed/framework.v1/facades/config"

	"github.com/RichardKnop/machinery/v2"
	redisBackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisBroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	configinstance "github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/gookit/color"

	"gopkg.in/go-mixed/framework.v1/contracts/event"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
)

func GetServer(connection string, queue string) (*machinery.Server, error) {
	if connection == "" {
		connection = config.GetString("queue.default")
	}

	driver := getDriver(connection)

	switch driver {
	case DriverSync:
		color.Yellowln("Queue sync driver doesn't need to be run")

		return nil, nil
	case DriverRedis:
		return getRedisServer(connection, queue), nil
	}

	return nil, fmt.Errorf("unknown queue driver: %s", driver)
}

func GetQueueName(connection, queue string) string {
	appName := config.GetString("app.name")
	if appName == "" {
		appName = "goravel"
	}
	if connection == "" {
		connection = config.GetString("queue.default")
	}
	if queue == "" {
		queue = config.GetString(fmt.Sprintf("queue.connections.%s.queue", connection), "default")
	}

	return fmt.Sprintf("%s_%s:%s", appName, "queues", queue)
}

func getDriver(connection string) string {
	if connection == "" {
		connection = config.GetString("queue.default")
	}

	return config.GetString(fmt.Sprintf("queue.connections.%s.driver", connection))
}

func getRedisServer(connection string, queue string) *machinery.Server {
	redisConfig, database, defaultQueue := getRedisConfig(connection)
	if queue == "" {
		queue = defaultQueue
	}

	cnf := &configinstance.Config{
		DefaultQueue: queue,
		Redis:        &configinstance.RedisConfig{},
	}

	broker := redisBroker.NewGR(cnf, []string{redisConfig}, database)
	backend := redisBackend.NewGR(cnf, []string{redisConfig}, database)
	lock := eager.New()

	return machinery.NewServer(cnf, broker, backend, lock)
}

func getRedisConfig(queueConnection string) (configResult string, database int, queue string) {
	connection := config.GetString(fmt.Sprintf("queue.connections.%s.connection", queueConnection))
	queue = GetQueueName(queueConnection, "")
	host := config.GetString(fmt.Sprintf("database.redis.%s.host", connection))
	password := config.GetString(fmt.Sprintf("database.redis.%s.password", connection))
	port := config.GetString(fmt.Sprintf("database.redis.%s.port", connection))
	database = config.GetInt(fmt.Sprintf("database.redis.%s.database", connection))

	if password == "" {
		configResult = host + ":" + port
	} else {
		configResult = password + "@" + host + ":" + port
	}

	return
}

func jobs2Tasks(jobs []queue.Job) (map[string]any, error) {
	tasks := make(map[string]any)

	for _, job := range jobs {
		if job.Signature() == "" {
			return nil, errors.New("the Signature of job can't be empty")
		}

		if tasks[job.Signature()] != nil {
			return nil, fmt.Errorf("job signature duplicate: %s, the names of Job and Listener cannot be duplicated", job.Signature())
		}

		tasks[job.Signature()] = job.Handle
	}

	return tasks, nil
}

func eventsToTasks(events map[event.Event][]event.Listener) (map[string]any, error) {
	tasks := make(map[string]any)

	for _, listeners := range events {
		for _, listener := range listeners {
			if listener.Signature() == "" {
				return nil, errors.New("the Signature of listener can't be empty")
			}

			if tasks[listener.Signature()] != nil {
				continue
			}

			tasks[listener.Signature()] = listener.Handle
		}
	}

	return tasks, nil
}
