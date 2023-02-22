package event

import (
	"gopkg.in/go-mixed/framework/contracts/event"
	"gopkg.in/go-mixed/framework/event/support"
)

type Application struct {
	events map[event.Event][]event.Listener
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Register(events map[event.Event][]event.Listener) {
	app.events = events
}

func (app *Application) GetEvents() map[event.Event][]event.Listener {
	return app.events
}

func (app *Application) Job(event event.Event, args []event.Arg) event.Task {
	return &support.Task{
		Events: app.events,
		Event:  event,
		Args:   args,
	}
}
