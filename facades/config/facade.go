package config

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/config"
	"time"
)

func getConfig() config.IConfig {
	return container.MustMake[config.IConfig]("config")
}

// Env Get config from env.
func Env(envName string, defaultValue ...any) any {
	return getConfig().Env(envName, defaultValue...)
}

// Add config to application.
func Add(name string, configuration map[string]any) {
	getConfig().Add(name, configuration)
}

// Get config from application.
func Get(path string, defaultValue ...any) any {
	return getConfig().Get(path, defaultValue...)
}

// GetString Get string type config from application.
func GetString(path string, defaultValue ...any) string {
	return getConfig().GetString(path, defaultValue...)
}

// GetInt Get int type config from application.
func GetInt(path string, defaultValue ...any) int {
	return getConfig().GetInt(path, defaultValue...)
}

// GetBool Get bool type config from application.
func GetBool(path string, defaultValue ...any) bool {
	return getConfig().GetBool(path, defaultValue...)
}

func GetMap(path string) map[string]any {
	return getConfig().GetMap(path)
}

func GetDuration(path string) time.Duration {
	return getConfig().GetDuration(path)
}

func GetTime(path string) time.Time {
	return getConfig().GetTime(path)
}
