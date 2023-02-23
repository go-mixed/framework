package config

import (
	"flag"
	"gopkg.in/go-mixed/framework.v1/container"
	contractsconfig "gopkg.in/go-mixed/framework.v1/contracts/config"
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

	container.Singleton((*Config)(nil), func(args ...any) (any, error) {
		return NewConfig(*env), nil
	})

	container.Alias("config", (*Config)(nil))
	container.Alias(contractsconfig.IConfig(nil), (*Config)(nil))
}

func (config *ServiceProvider) Boot() {

}
