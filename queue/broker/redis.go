package broker

import (
	"fmt"
	"github.com/RichardKnop/machinery/v2"
	redisBackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisBroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	configinstance "github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/locks/eager"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
	"gopkg.in/go-mixed/framework.v1/facades/config"
)

type RedisBroker struct {
	machineryBroker
}

var _ queue.IBroker = (*RedisBroker)(nil)

func NewRedisBroker(connectionName string, queueName string) (*RedisBroker, error) {
	brokerUrl, database, defaultQueueName := getRedisConfig(connectionName)
	if queueName == "" {
		queueName = defaultQueueName
	}

	cnf := &configinstance.Config{
		//Broker:          brokerUrl,
		DefaultQueue:    queueName,
		Redis:           &configinstance.RedisConfig{},
		ResultsExpireIn: int(config.GetDuration("queue.result_expire").Seconds()),
	}

	broker := redisBroker.NewGR(cnf, []string{brokerUrl}, database)
	backend := redisBackend.NewGR(cnf, []string{brokerUrl}, database)
	lock := eager.New()

	b := &RedisBroker{
		machineryBroker{
			server:           machinery.NewServer(cnf, broker, backend, lock),
			jobMap:           container.MustMakeAs("queue.job_map", queue.IJobMap(nil)),
			defaultQueueName: queueName,
			connectionName:   connectionName,
		},
	}
	if err := b.registerTasks(); err != nil {
		return nil, err
	}

	return b, nil
}

func getRedisConfig(queueConnection string) (brokerUrl string, database int, queue string) {
	connection := config.GetString(fmt.Sprintf("queue.connections.%s.connection", queueConnection))
	queue = GetQueueName(queueConnection, "")

	keyPrefix := fmt.Sprintf("database.redis.%s.", connection)
	host := config.GetString(keyPrefix + "host")
	password := config.GetString(keyPrefix + "password")
	port := config.GetInt(keyPrefix+"port", 6379)
	database = config.GetInt(keyPrefix+"database", 0)

	if password == "" {
		brokerUrl = fmt.Sprintf("%s:%d", host, port)
	} else {
		brokerUrl = fmt.Sprintf("%s@%s:%d", password, host, port)
	}

	return
}
