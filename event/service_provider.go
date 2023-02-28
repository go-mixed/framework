package event

import (
	"gopkg.in/go-mixed/framework.v1/contracts/console"
	eventConsole "gopkg.in/go-mixed/framework.v1/event/console"
	"gopkg.in/go-mixed/framework.v1/facades"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register() {
	facades.Event = NewApplication()
}

func (receiver *ServiceProvider) Boot() {
	receiver.registerCommands()
}

func (receiver *ServiceProvider) registerCommands() {
	artisan.Register([]console.Command{
		&eventConsole.EventMakeCommand{},
		&eventConsole.ListenerMakeCommand{},
	})
}
