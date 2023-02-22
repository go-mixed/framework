package console

import (
	"gopkg.in/go-mixed/framework/console/console"
	console2 "gopkg.in/go-mixed/framework/contracts/console"
	"gopkg.in/go-mixed/framework/facades"
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
