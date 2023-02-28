package validation

import (
	consolecontract "gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/facades"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
	"gopkg.in/go-mixed/framework.v1/validation/console"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Validation = NewValidation()
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	artisan.Register([]consolecontract.Command{
		&console.RuleMakeCommand{},
	})
}
