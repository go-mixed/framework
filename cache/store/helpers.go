package store

import (
	"gopkg.in/go-mixed/framework.v1/facades/config"
)

func prefix() string {
	return config.GetString("cache.prefix") + ":"
}
