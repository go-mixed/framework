package event

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/event"
)

func getEvent() event.IInstance {
	return container.MustMake[event.IInstance]("instance")
}

func Register(event map[event.Event][]event.Listener) {
	getEvent().Register(event)
}

func Job(event event.Event, args []event.Arg) event.Task {
	return getEvent().Job(event, args)
}

func GetEvents() map[event.Event][]event.Listener {
	return getEvent().GetEvents()
}
