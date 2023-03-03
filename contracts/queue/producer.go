package queue

import "time"

//go:generate mockery --name=IProducer
type IProducer interface {
	AddJob(job IBrokerJob) IProducer
	Dispatch() error
	DispatchSync() error
	Delay(eta time.Time) IProducer

	OnConnection(connection string) IProducer
	OnQueue(queue string) IProducer
}
