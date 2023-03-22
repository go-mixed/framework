package broker

import (
	"go.uber.org/multierr"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/manager"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
)

type SyncBroker struct {
	jobMap queue.IJobMap
}

var _ queue.IBroker = (*SyncBroker)(nil)

func NewSyncBroker(connection string) *SyncBroker {
	return &SyncBroker{
		jobMap: container.MustMakeAs("queue.job_map", queue.IJobMap(nil)),
	}
}

func (s *SyncBroker) Connection(name string) queue.IBroker {
	return container.MustMakeAs("queue.manager", manager.IManager[queue.IBroker](nil)).MustDriver(name)
}

func (s *SyncBroker) RunServe(queueName string, concurrentCount int) error {
	return nil
}

func (s *SyncBroker) AddJob(jobs ...queue.IBrokerJob) error {
	var err error
	for _, job := range jobs {
		err = multierr.Append(err, s.jobMap.Invoke(job.Signature(), decodeArgs(job.Arguments())...))
	}

	return err
}

func (s *SyncBroker) AddChainJobs(jobs ...queue.IBrokerJob) error {
	return s.AddJob(jobs...)
}
