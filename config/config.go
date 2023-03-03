package config

import (
	"gopkg.in/go-mixed/framework.v1/contracts/config"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"gopkg.in/go-mixed/framework.v1/support/file"
	"gopkg.in/go-mixed/framework.v1/testing"
)

type Config struct {
	vip *viper.Viper
}

var _ config.IConfig = (*Config)(nil)

func NewConfig(envPath string) *Config {
	if !file.Exists(envPath) {
		color.Redln("Please create .env and initialize it first.")
		color.Warnln("Run command: \ncp .env.example .env && go run . artisan key:generate")
		os.Exit(0)
	}

	conf := &Config{}
	conf.vip = viper.New()
	conf.vip.SetConfigName(envPath)
	conf.vip.SetConfigType("env")
	conf.vip.AddConfigPath(".")
	err := conf.vip.ReadInConfig()
	if err != nil {
		if !testing.RunInTest() {
			panic(err.Error())
		}
	}
	conf.vip.SetEnvPrefix("laravel")
	conf.vip.AutomaticEnv()

	return conf
}

// Env Get config from env.
func (conf *Config) Env(envName string, defaultValue ...any) any {
	value := conf.Get(envName, defaultValue...)
	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}

		return nil
	}

	return value
}

// Add config to application.
func (conf *Config) Add(name string, configuration map[string]any) {
	conf.vip.Set(name, configuration)
}

// Get config from application.
func (conf *Config) Get(path string, defaultValue ...any) any {
	if !conf.vip.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}

	return conf.vip.Get(path)
}

// GetString Get string type config from application.
func (conf *Config) GetString(path string, defaultValue ...any) string {
	value := cast.ToString(conf.Get(path, defaultValue...))
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0].(string)
		}

		return ""
	}

	return value
}

// GetInt Get int type config from application.
func (conf *Config) GetInt(path string, defaultValue ...any) int {
	value := conf.Get(path, defaultValue...)
	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0].(int)
		}

		return 0
	}

	return cast.ToInt(value)
}

// GetBool Get bool type config from application.
func (conf *Config) GetBool(path string, defaultValue ...any) bool {
	value := conf.Get(path, defaultValue...)
	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0].(bool)
		}

		return false
	}

	return cast.ToBool(value)
}

func (conf *Config) GetMap(path string) map[string]any {
	value := conf.Get(path)
	actual := cast.ToStringMap(value)
	if actual != nil {
		return actual
	}

	return map[string]any{}
}
