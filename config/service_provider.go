package config

import (
	"flag"
	"gopkg.in/go-mixed/framework.v1/facades"
	"gopkg.in/go-mixed/framework.v1/testing"
)

type ServiceProvider struct {
}

func (config *ServiceProvider) Register() {
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

func (config *ServiceProvider) Boot() {

}
