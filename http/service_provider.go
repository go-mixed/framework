package http

import (
	consolecontract "gopkg.in/go-mixed/framework/contracts/console"
	"gopkg.in/go-mixed/framework/facades"
	"gopkg.in/go-mixed/framework/http/console"
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
