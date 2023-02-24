package database

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/database/orm"

	consolecontract "gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/database/console"
	"gopkg.in/go-mixed/framework.v1/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	container.Singleton((*ConnectionManager)(nil), func(args ...any) (any, error) {
		manager := NewConnectionManager()
		manager.extendConnections()
		return manager, nil
	})

	container.Alias("database.manager", (*ConnectionManager)(nil))
	container.Singleton((orm.IOrm)(nil), func(args ...any) (any, error) {
		return container.MustMake[*ConnectionManager]((*ConnectionManager)(nil)).DefaultDriver()
	})
	container.Alias((*Orm)(nil), (orm.IOrm)(nil))
	container.Alias("database", (orm.IOrm)(nil))
	container.Alias("db", (orm.IOrm)(nil))
	container.Alias("orm", (orm.IOrm)(nil))

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
