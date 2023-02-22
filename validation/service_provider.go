package validation

import (
	consolecontract "gopkg.in/go-mixed/framework/contracts/console"
	"gopkg.in/go-mixed/framework/facades"
	"gopkg.in/go-mixed/framework/validation/console"
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
	facades.Artisan.Register([]consolecontract.Command{
		&console.RuleMakeCommand{},
	})
}
