package config

import (
	"flag"
	"gopkg.in/go-mixed/framework/contracts/container"

	"gopkg.in/go-mixed/framework/facades"
	"gopkg.in/go-mixed/framework/testing"
)

type ServiceProvider struct {
}

func (config *ServiceProvider) Register(container container.IContainer) {
	var env *string
	if testing.RunInTest() {
		testEnv := ".env"
		env = &testEnv
	} else {
		env = flag.String("env", ".env", "custom .env path")
		flag.Parse()
	}
	facades.Config = NewModule(*env)
}

func (config *ServiceProvider) Boot(container container.IContainer) {

}
