package console

import (
	"gopkg.in/go-mixed/framework.v1/console/console"
	console2 "gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/facades"
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register() {
	facades.Artisan = NewApplication()
}

func (receiver *ServiceProvider) Boot() {
	receiver.registerCommands()
}

func (receiver *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]console2.Command{
		&console.ListCommand{},
		&console.KeyGenerateCommand{},
		&console.MakeCommand{},
	})
}
