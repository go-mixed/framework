package queue

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
	queue2 "gopkg.in/go-mixed/framework.v1/queue"
)

func getJobMap() queue.IJobMap {
	return container.MustMakeAs("queue.job_map", queue.IJobMap(nil))
}

func getProducer() queue.IProducer {
	return container.MustMakeAs("queue.producer", queue.IProducer(nil))
}

func getManager() *queue2.QueueManager {
	return container.MustMakeAs("queue.manager", (*queue2.QueueManager)(nil))
}

func Register(jobs ...queue.IJob) queue.IJobMap {
	return getJobMap().Register(jobs...)
}

// RegisterWithName job-func map
func RegisterWithName(name string, jobFunc queue.JobFunc) queue.IJobMap {
	return getJobMap().RegisterWithName(name, jobFunc)
}

func Registers(jobMap map[string]queue.JobFunc) queue.IJobMap {
	return getJobMap().Registers(jobMap)
}

func JobByName[T queue.Argument | any](name string, args ...T) queue.IProducer {
	return getProducer().AddJob(queue.MakeJobWithName[T](name, args...))
}

func Job[T queue.Argument | any](job queue.IJob, args ...T) queue.IProducer {
	return getProducer().AddJob(queue.MakeJob[T](job, args...))
}

// Dispatch invoke a job asynchronous
func Dispatch(name string, args ...any) error {
	return getProducer().AddJob(queue.MakeJobWithName(name, args...)).Dispatch()
}

// DispatchSync invoke a job synchronous
func DispatchSync(name string, args ...any) error {
	return getProducer().AddJob(queue.MakeJobWithName(name, args...)).DispatchSync()
}

// RunServe run a queue server
func RunServe(connectionName string, queueName string, concurrentCount int) error {
	return getManager().Connection(connectionName).RunServe(queueName, concurrentCount)
}
