package auth

import (
	"context"

	"gopkg.in/go-mixed/framework/auth/access"
	"gopkg.in/go-mixed/framework/auth/console"
	contractconsole "gopkg.in/go-mixed/framework/contracts/console"
	"gopkg.in/go-mixed/framework/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Auth = NewAuth(facades.Config.GetString("auth.defaults.guard"))
	facades.Gate = access.NewGate(context.Background())
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]contractconsole.Command{
		&console.JwtSecretCommand{},
		&console.PolicyMakeCommand{},
	})
}
