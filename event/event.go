package event

import (
	"gopkg.in/go-mixed/framework.v1/contracts/event"
	"gopkg.in/go-mixed/framework.v1/event/support"
)

type Event struct {
	events map[event.Event][]event.Listener
}

func NewEvent() *Event {
	return &Event{}
}

func (e *Event) Register(events map[event.Event][]event.Listener) {
	e.events = events
}

func (e *Event) GetEvents() map[event.Event][]event.Listener {
	return e.events
}

func (e *Event) Job(event event.Event, args []event.Arg) event.Task {
	return &support.Task{
		Events: e.events,
		Event:  event,
		Args:   args,
	}
}
