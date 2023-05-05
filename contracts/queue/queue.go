package queue

//go:generate mockery --name=IBroker
type IBroker interface {
	IConsumer

	Connection(name string, queue string) IBroker
	AddJob(jobs ...IBrokerJob) error
	AddChainJobs(jobs ...IBrokerJob) error
}

type IConsumer interface {
	RunServe(queueName string, concurrentCount int) error
}

type Argument struct {
	Name  string
	Type  string
	Value any
}
