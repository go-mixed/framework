package broker

import (
	"fmt"
	"github.com/RichardKnop/machinery/v2"
	redisBackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisBroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	configinstance "github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/RichardKnop/machinery/v2/tasks"
	"go.uber.org/multierr"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	ievent "gopkg.in/go-mixed/framework.v1/facades/event"
	"runtime"
)

type RedisBroker struct {
	server *machinery.Server
	jobMap queue.IJobMap

	defaultQueueName string
}

var _ queue.IBroker = (*RedisBroker)(nil)

func NewRedisBroker(connectionName string) (*RedisBroker, error) {
	redisConfig, database, queueName := getRedisConfig(connectionName)

	cnf := &configinstance.Config{
		DefaultQueue: queueName,
		Redis:        &configinstance.RedisConfig{},
	}

	broker := redisBroker.NewGR(cnf, []string{redisConfig}, database)
	backend := redisBackend.NewGR(cnf, []string{redisConfig}, database)
	lock := eager.New()

	b := &RedisBroker{
		server:           machinery.NewServer(cnf, broker, backend, lock),
		jobMap:           container.MustMake[queue.IJobMap]("queue.job_map"),
		defaultQueueName: queueName,
	}
	if err := b.registerTasks(); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *RedisBroker) registerTasks() error {
	var err error

	// register the events
	eventTasks, err := eventsToJobMap(ievent.GetEvents())
	if err != nil {
		return err
	}
	if err = b.server.RegisterTasks(eventTasks); err != nil {
		return err
	}

	// register the custom
	for name, fn := range b.jobMap.GetMap() {
		err = multierr.Append(err, b.server.RegisterTask(name, fn))
	}
	return err
}

func (b *RedisBroker) AddJob(jobs ...queue.IBrokerJob) error {
	var err error
	for _, job := range jobs {
		err = multierr.Append(err, b.sendToTask(job))
	}
	return err
}

func (b *RedisBroker) AddChainJobs(jobs ...queue.IBrokerJob) error {
	//TODO implement me
	panic("implement me")
}

func (b *RedisBroker) sendToTask(job queue.IBrokerJob) error {

	_, err := b.server.SendTask(&tasks.Signature{
		Name: job.Signature(),
		Args: encodeArgs(job.Arguments()),
		ETA:  job.ETA(),
	})

	return err
}

func (b *RedisBroker) RunServe(queueName string, concurrentCount int) error {
	if concurrentCount <= 0 {
		concurrentCount = runtime.NumCPU()
	}
	if queueName == "" {
		queueName = b.defaultQueueName
	}
	worker := b.server.NewWorker(queueName, concurrentCount)
	return worker.Launch()
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
