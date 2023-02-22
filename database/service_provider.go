package database

import (
	"context"

	consolecontract "gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/database/console"
	"gopkg.in/go-mixed/framework.v1/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Orm = NewOrm(context.Background())
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]consolecontract.Command{
		&console.MigrateMakeCommand{},
		&console.MigrateCommand{},
		&console.MigrateRollbackCommand{},
		&console.ModelMakeCommand{},
	})
}
