package config

import "time"

//go:generate mockery --name=Config
type IConfig interface {
	//Env Get config from env.
	Env(envName string, defaultValue ...any) any
	//Add config to application.
	Add(name string, configuration map[string]any)
	//Get config from application.
	Get(path string, defaultValue ...any) any
	//GetString Get string type config from application.
	GetString(path string, defaultValue ...any) string
	//GetInt Get int type config from application.
	GetInt(path string, defaultValue ...any) int
	//GetBool Get bool type config from application.
	GetBool(path string, defaultValue ...any) bool

	GetMap(path string) map[string]any

	GetDuration(path string) time.Duration
	GetTime(path string) time.Time
}
