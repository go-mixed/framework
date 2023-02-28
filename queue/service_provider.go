package queue

import (
	"gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/facades"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
	queueConsole "gopkg.in/go-mixed/framework.v1/queue/console"
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register() {
	facades.Queue = NewApplication()
}

func (receiver *ServiceProvider) Boot() {
	receiver.registerCommands()
}

func (receiver *ServiceProvider) registerCommands() {
	artisan.Register([]console.Command{
		&queueConsole.JobMakeCommand{},
	})
}
