package console

import (
	"gopkg.in/go-mixed/framework.v1/console/console"
	"gopkg.in/go-mixed/framework.v1/container"
	console2 "gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
)

type ServiceProvider struct {
}

func (sp *ServiceProvider) Register() {
	container.Singleton((*Console)(nil), func(args ...any) (any, error) {
		return NewConsole(), nil
	})
	container.Alias("artisan", (*Console)(nil))
}

func (sp *ServiceProvider) Boot() {
	sp.registerCommands()
}

func (sp *ServiceProvider) registerCommands() {
	artisan.Register([]console2.ICommand{
		&console.ListCommand{},
		&console.KeyGenerateCommand{},
		&console.MakeCommand{},
	})
}
