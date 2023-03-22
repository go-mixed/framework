package queue

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
	"gopkg.in/go-mixed/framework.v1/queue/broker"
	queueConsole "gopkg.in/go-mixed/framework.v1/queue/console"
	"gopkg.in/go-mixed/framework.v1/queue/register"
)

type ServiceProvider struct {
}

func (sp *ServiceProvider) Register() {

	// queue manager
	container.Singleton((*QueueManager)(nil), func(args ...any) (any, error) {
		m := NewQueueManager()
		m.Extend("sync", func(driverName string, args ...any) (queue.IBroker, error) {
			return broker.NewSyncBroker(driverName), nil
		})
		m.Extend("redis", func(driverName string, args ...any) (queue.IBroker, error) {
			return broker.NewRedisBroker(driverName)
		})
		m.Extend("amqp", func(driverName string, args ...any) (queue.IBroker, error) {
			return broker.NewAmqpBroker(driverName)
		})

		return m, nil
	})
	container.Alias("queue.manager", (*QueueManager)(nil))

	// job registered map
	container.Singleton((queue.IJobMap)(nil), func(args ...any) (any, error) {
		return register.NewJobMap(), nil
	})
	container.Alias("queue.job_map", (queue.IJobMap)(nil))
	container.Alias((*register.JobMap)(nil), (queue.IJobMap)(nil))

	// default broker
	container.Singleton(queue.IBroker(nil), func(args ...any) (any, error) {
		return container.MustMakeAs("queue.manager", (*QueueManager)(nil)).DefaultDriver()
	})
	container.Alias("queue", queue.IBroker(nil))
	container.Alias("queue.connection", queue.IBroker(nil))

	// no-singleton producer
	container.Bind(queue.IProducer(nil), func(args ...any) (any, error) {
		return NewProducer(), nil
	}, false)
	container.Alias("queue.producer", queue.IProducer(nil))
}

func (sp *ServiceProvider) Boot() {
	sp.registerCommands()
}

func (sp *ServiceProvider) registerCommands() {
	artisan.Register([]console.Command{
		&queueConsole.JobMakeCommand{},
		&queueConsole.QueueWorkCommand{},
	})
}
