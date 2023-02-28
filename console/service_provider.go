package console

import (
	"gopkg.in/go-mixed/framework.v1/console/console"
	"gopkg.in/go-mixed/framework.v1/container"
	console2 "gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register() {
	container.Singleton((*Application)(nil), func(args ...any) (any, error) {
		return NewApplication(), nil
	})
	container.Alias("artisan", (*Application)(nil))
}

func (receiver *ServiceProvider) Boot() {
	receiver.registerCommands()
}

func (receiver *ServiceProvider) registerCommands() {
	artisan.Register([]console2.Command{
		&console.ListCommand{},
		&console.KeyGenerateCommand{},
		&console.MakeCommand{},
	})
}
