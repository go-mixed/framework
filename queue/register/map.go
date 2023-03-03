package register

import (
	"errors"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
)

type JobMap struct {
	bindings map[string]queue.JobFunc
}

var _ queue.IJobMap = (*JobMap)(nil)

var ErrJobNotFound = errors.New("the job is undefined with this name")

func NewJobMap() *JobMap {
	return &JobMap{bindings: make(map[string]queue.JobFunc)}
}

func (m *JobMap) Register(jobs ...queue.IJob) queue.IJobMap {
	for _, job := range jobs {
		m.bindings[job.Signature()] = job.Handle
	}
	return m
}

func (m *JobMap) RegisterWithName(name string, fn queue.JobFunc) queue.IJobMap {
	m.bindings[name] = fn
	return m
}

func (m *JobMap) Registers(fns map[string]queue.JobFunc) queue.IJobMap {
	for name, fn := range fns {
		m.bindings[name] = fn
	}
	return m
}

func (m *JobMap) Get(name string) queue.JobFunc {
	return m.bindings[name]
}

func (m *JobMap) Invoke(name string, args ...any) error {
	fn, ok := m.bindings[name]
	if !ok {
		return ErrJobNotFound
	}

	return fn(args...)
}

func (m *JobMap) GetMap() map[string]queue.JobFunc {
	return m.bindings
}
