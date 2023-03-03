package queue

//go:generate mockery --name=IBroker
type IBroker interface {
	IConsumer
	AddJob(jobs ...IBrokerJob) error
	AddChainJobs(jobs ...IBrokerJob) error
}

type IConsumer interface {
	RunServe(queueName string, concurrentCount int) error
}

type Argument struct {
	Type  string
	Value any
}
