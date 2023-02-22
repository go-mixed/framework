package console

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gookit/color"

	"gopkg.in/go-mixed/framework/contracts/console"
	"gopkg.in/go-mixed/framework/contracts/console/command"
)

type MigrateRollbackCommand struct {
}

// Signature The name and signature of the console command.
func (receiver *MigrateRollbackCommand) Signature() string {
	return "migrate:rollback"
}

// Description The console command description.
func (receiver *MigrateRollbackCommand) Description() string {
	return "Rollback the database migrations"
}

// Extend The console command extend.
func (receiver *MigrateRollbackCommand) Extend() command.Extend {
	return command.Extend{
		Category: "migrate",
		Flags: []command.Flag{
			{
				Name:  "step",
				Value: "1",
				Usage: "rollback steps",
			},
		},
	}
}

// Handle Execute the console command.
func (receiver *MigrateRollbackCommand) Handle(ctx console.Context) error {
	m, err := getMigrate()
	if err != nil {
		return err
	}
	if m == nil {
		color.Yellowln("Please fill database config first")

		return nil
	}

	stepString := "-" + ctx.Option("step")
	step, err := strconv.Atoi(stepString)
	if err != nil {
		color.Redln("Migration failed: invalid step", ctx.Option("step"))

		return nil
	}

	if err := m.Steps(step); err != nil && err != migrate.ErrNoChange && err != migrate.ErrNilVersion {
		switch err.(type) {
		case migrate.ErrShortLimit:
		default:
			color.Redln("Migration failed:", err.Error())

			return nil
		}
	}

	color.Greenln("Migration rollback success")

	return nil
}
