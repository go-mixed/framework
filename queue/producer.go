package queue

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/queue/broker"
	"time"

	"gopkg.in/go-mixed/framework.v1/contracts/queue"
)

type Producer struct {
	Chain      bool
	Jobs       []queue.IBrokerJob
	connection string
	queue      string
	eta        *time.Time
}

var _ queue.IProducer = (*Producer)(nil)

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) AddJob(job queue.IBrokerJob) queue.IProducer {
	p.Jobs = append(p.Jobs, job)
	return p
}

func (p *Producer) Dispatch() error {
	b := container.MustMakeAs("queue.manager", (*QueueManager)(nil)).Connection(p.connection)
	return b.AddJob(p.Jobs...)
}

func (p *Producer) DispatchSync() error {
	b := container.MustMakeAs("queue.manager", (*QueueManager)(nil)).Connection("sync")
	return b.AddJob(p.Jobs...)
}

func (p *Producer) OnConnection(connection string) queue.IProducer {
	p.connection = connection
	return p
}

func (p *Producer) OnQueue(queue string) queue.IProducer {
	p.queue = broker.GetQueueName(p.connection, queue)
	return p
}

func (p *Producer) Delay(eta time.Time) queue.IProducer {
	p.eta = &eta
	return p
}
