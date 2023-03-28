package broker

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/tasks"
	"go.uber.org/multierr"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/manager"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
	ievent "gopkg.in/go-mixed/framework.v1/facades/event"
	"runtime"
)

type machineryBroker struct {
	server *machinery.Server
	jobMap queue.IJobMap

	connectionName   string
	defaultQueueName string
}

func (b *machineryBroker) Connection(name string) queue.IBroker {
	return container.MustMakeAs("queue.manager", manager.IManager[queue.IBroker](nil)).MustDriver(name)
}

func (b *machineryBroker) registerTasks() error {
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

func (b *machineryBroker) AddJob(jobs ...queue.IBrokerJob) error {
	var err error
	for _, job := range jobs {
		err = multierr.Append(err, b.sendToTask(job))
	}
	return err
}

func (b *machineryBroker) AddChainJobs(jobs ...queue.IBrokerJob) error {
	var signatures []*tasks.Signature
	for _, job := range jobs {
		signatures = append(signatures, &tasks.Signature{
			Name: job.Signature(),
			Args: encodeArgs(job.Arguments()),
			ETA:  job.ETA(),
		})
	}
	chain, err := tasks.NewChain(signatures...)
	if err != nil {
		return err
	}

	_, err = b.server.SendChain(chain)
	return err
}

func (b *machineryBroker) sendToTask(job queue.IBrokerJob) error {

	_, err := b.server.SendTask(&tasks.Signature{
		Name: job.Signature(),
		Args: encodeArgs(job.Arguments()),
		ETA:  job.ETA(),
	})

	return err
}

func (b *machineryBroker) RunServe(queueName string, concurrentCount int) error {
	if concurrentCount <= 0 {
		concurrentCount = runtime.NumCPU()
	}
	if queueName == "" {
		queueName = b.defaultQueueName
	} else {
		queueName = GetQueueName(b.connectionName, queueName)
	}
	worker := b.server.NewWorker(queueName, concurrentCount)
	return worker.Launch()
}
