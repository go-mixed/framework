package validation

import (
	"gopkg.in/go-mixed/framework.v1/container"
	consolecontract "gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/contracts/validation"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
	"gopkg.in/go-mixed/framework.v1/validation/console"
)

type ServiceProvider struct {
}

func (sp *ServiceProvider) Register() {
	container.Singleton((validation.IValidation)(nil), func(args ...any) (any, error) {
		return NewValidation(), nil
	})
	container.Alias("validation", (validation.IValidation)(nil))
}

func (sp *ServiceProvider) Boot() {
	sp.registerCommands()
}

func (sp *ServiceProvider) registerCommands() {
	artisan.Register([]consolecontract.Command{
		&console.RuleMakeCommand{},
	})
}
