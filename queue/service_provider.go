package queue

import (
	"gopkg.in/go-mixed/framework/contracts/console"
	"gopkg.in/go-mixed/framework/facades"
	queueConsole "gopkg.in/go-mixed/framework/queue/console"
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
	facades.Artisan.Register([]console.Command{
		&queueConsole.JobMakeCommand{},
	})
}
