package http

import (
	consolecontract "gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/facades"
	"gopkg.in/go-mixed/framework.v1/http/console"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.RateLimiter = NewRateLimiter()
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]consolecontract.Command{
		&console.RequestMakeCommand{},
		&console.ControllerMakeCommand{},
		&console.MiddlewareMakeCommand{},
	})
}
