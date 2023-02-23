package config

import (
	"gopkg.in/go-mixed/framework.v1/contracts/config"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"gopkg.in/go-mixed/framework.v1/support"
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

	app := &Config{}
	app.vip = viper.New()
	app.vip.SetConfigName(envPath)
	app.vip.SetConfigType("env")
	app.vip.AddConfigPath(".")
	err := app.vip.ReadInConfig()
	if err != nil {
		if !testing.RunInTest() {
			panic(err.Error())
		}
	}
	app.vip.SetEnvPrefix("goravel")
	app.vip.AutomaticEnv()

	appKey := app.Env("APP_KEY")
	if appKey == nil && support.Env != support.EnvArtisan {
		color.Redln("Please initialize APP_KEY first.")
		color.Warnln("Run command: \ngo run . artisan key:generate")
		os.Exit(0)
	}

	if len(appKey.(string)) != 32 {
		color.Redln("Invalid APP_KEY, please reset it.")
		color.Warnln("Run command: \ngo run . artisan key:generate")
		os.Exit(0)
	}

	return app
}

// Env Get config from env.
func (app *Config) Env(envName string, defaultValue ...any) any {
	value := app.Get(envName, defaultValue...)
	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}

		return nil
	}

	return value
}

// Add config to application.
func (app *Config) Add(name string, configuration map[string]any) {
	app.vip.Set(name, configuration)
}

// Get config from application.
func (app *Config) Get(path string, defaultValue ...any) any {
	if !app.vip.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}

	return app.vip.Get(path)
}

// GetString Get string type config from application.
func (app *Config) GetString(path string, defaultValue ...any) string {
	value := cast.ToString(app.Get(path, defaultValue...))
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0].(string)
		}

		return ""
	}

	return value
}

// GetInt Get int type config from application.
func (app *Config) GetInt(path string, defaultValue ...any) int {
	value := app.Get(path, defaultValue...)
	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0].(int)
		}

		return 0
	}

	return cast.ToInt(value)
}

// GetBool Get bool type config from application.
func (app *Config) GetBool(path string, defaultValue ...any) bool {
	value := app.Get(path, defaultValue...)
	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0].(bool)
		}

		return false
	}

	return cast.ToBool(value)
}

func (app *Config) GetMap(path string) map[string]any {
	value := app.Get(path)
	actual := cast.ToStringMap(value)
	if actual != nil {
		return actual
	}

	return map[string]any{}
}
