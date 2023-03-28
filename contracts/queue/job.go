package queue

import (
	"reflect"
	"time"
)

type JobFunc func(args ...any) error

//go:generate mockery --name=IJob
type IJob interface {
	Signature() string
	Handle(args ...any) error
}

//go:generate mockery --name=IBrokerJob
type IBrokerJob interface {
	Signature() string
	Arguments() []Argument
	ETA() *time.Time
}

//go:generate mockery --name=IJobMap
type IJobMap interface {
	Register(jobs ...IJob) IJobMap
	RegisterWithName(name string, fn JobFunc) IJobMap
	Registers(map[string]JobFunc) IJobMap
	Get(name string) JobFunc

	Invoke(name string, args ...any) error
	GetMap() map[string]JobFunc
}

type Job struct {
	signature string
	args      []Argument
	eta       time.Time
}

func (j *Job) ETA() *time.Time {
	if j.eta.IsZero() {
		return nil
	}
	return &j.eta
}

func (j *Job) Arguments() []Argument {
	return j.args
}

func (j *Job) Signature() string {
	return j.signature
}

func MakeJob[T Argument | any](job IJob, args ...T) IBrokerJob {
	return MakeJobWithName[T](job.Signature(), args...)
}

func MakeJobWithName[T Argument | any](name string, args ...T) IBrokerJob {
	var mArgs []Argument
	for _, arg := range args {
		if _args, ok := any(arg).(Argument); ok {
			mArgs = append(mArgs, _args)
		} else {
			mArgs = append(mArgs, Argument{
				Type:  reflect.TypeOf(arg).Name(),
				Value: arg,
			})
		}
	}
	return &Job{
		signature: name,
		args:      mArgs,
	}
}
