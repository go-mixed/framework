package auth

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/auth"
	accesscontract "gopkg.in/go-mixed/framework.v1/contracts/auth/access"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
	"gopkg.in/go-mixed/framework.v1/facades/config"

	"gopkg.in/go-mixed/framework.v1/auth/access"
	"gopkg.in/go-mixed/framework.v1/auth/console"
	contractconsole "gopkg.in/go-mixed/framework.v1/contracts/console"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	container.Singleton((*Auth)(nil), func(args ...any) (any, error) {
		return NewAuth(config.GetString("auth.defaults.guard")), nil
	})
	container.Alias("auth", (*Auth)(nil))
	container.Alias(auth.IAuth(nil), (*Auth)(nil))

	container.Singleton((*access.Gate)(nil), func(args ...any) (any, error) {
		return access.NewGate(context.Background()), nil
	})
	container.Alias("gate", (*access.Gate)(nil))
	container.Alias(accesscontract.IGate(nil), (*access.Gate)(nil))
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	artisan.Register([]contractconsole.Command{
		&console.JwtSecretCommand{},
		&console.PolicyMakeCommand{},
	})
}
