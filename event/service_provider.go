package event

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/contracts/event"
	eventConsole "gopkg.in/go-mixed/framework.v1/event/console"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register() {
	container.Singleton((event.IInstance)(nil), func(args ...any) (any, error) {
		return NewEvent(), nil
	})
	container.Alias("event", (event.IInstance)(nil))
}

func (receiver *ServiceProvider) Boot() {
	receiver.registerCommands()
}

func (receiver *ServiceProvider) registerCommands() {
	artisan.Register([]console.ICommand{
		&eventConsole.EventMakeCommand{},
		&eventConsole.ListenerMakeCommand{},
	})
}
