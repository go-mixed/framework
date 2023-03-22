package support

import (
	"fmt"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/event"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
)

type Task struct {
	Events      map[event.Event][]event.Listener
	Event       event.Event
	Args        []event.Arg
	handledArgs []event.Arg
	mapArgs     []any
}

func (t *Task) Dispatch() error {
	listeners, ok := t.Events[t.Event]
	if !ok {
		return fmt.Errorf("event not found: %v", t.Event)
	}

	handledArgs, err := t.Event.Handle(t.Args)
	if err != nil {
		return err
	}

	t.handledArgs = handledArgs

	var mapArgs []any
	for _, arg := range t.handledArgs {
		mapArgs = append(mapArgs, arg.Value)
	}
	t.mapArgs = mapArgs

	for _, listener := range listeners {
		var err error
		queue := listener.Queue(t.mapArgs...)
		if queue.Enable {
			err = t.dispatchAsync(listener)
		} else {
			err = t.dispatchSync(listener)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Task) dispatchAsync(listener event.Listener) error {
	b := container.MustMakeAs("queue.connection", queue.IBroker(nil))

	var args []queue.Argument
	for _, arg := range t.handledArgs {
		args = append(args, queue.Argument{
			Type:  arg.Type,
			Value: arg.Value,
		})
	}

	if err := b.AddJob(queue.MakeJobWithName(listener.Signature(), args...)); err != nil {
		return err
	}

	return nil
}

func (t *Task) dispatchSync(listen event.Listener) error {
	return listen.Handle(t.mapArgs...)
}
