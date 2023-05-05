package queue

import (
	"gopkg.in/go-mixed/framework.v1/container"
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
	driver := p.connection
	if p.queue != "" {
		driver = p.connection + "|" + p.queue
	}
	b := container.MustMakeAs("queue.manager", (*QueueManager)(nil)).Connection(driver)
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
	p.queue = queue
	return p
}

func (p *Producer) Delay(eta time.Time) queue.IProducer {
	p.eta = &eta
	return p
}
