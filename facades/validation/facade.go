package validation

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/validation"
)

func getValidation() validation.IValidation {
	return container.MustMakeAs("validation", validation.IValidation(nil))
}

func Make(data any, rules map[string]string, options ...validation.Option) (validation.Validator, error) {
	return getValidation().Make(data, rules, options...)
}

func AddRules(rules []validation.Rule) error {
	return getValidation().AddRules(rules)
}

func Rules() []validation.Rule {
	return getValidation().Rules()
}
