package http

import (
	"gopkg.in/go-mixed/framework.v1/container"
	consolecontract "gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
	"gopkg.in/go-mixed/framework.v1/http/console"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	container.Singleton((*RateLimiter)(nil), func(args ...any) (any, error) {
		return NewRateLimiter(), nil
	})
	container.Alias("http", (*RateLimiter)(nil))
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	artisan.Register([]consolecontract.Command{
		&console.RequestMakeCommand{},
		&console.ControllerMakeCommand{},
		&console.MiddlewareMakeCommand{},
	})
}
