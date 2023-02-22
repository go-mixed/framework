package database

import (
	"context"

	consolecontract "gopkg.in/go-mixed/framework/contracts/console"
	"gopkg.in/go-mixed/framework/database/console"
	"gopkg.in/go-mixed/framework/facades"
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
