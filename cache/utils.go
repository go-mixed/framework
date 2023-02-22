package cache

import (
	"gopkg.in/go-mixed/framework/facades"
)

func prefix() string {
	return facades.Config.GetString("cache.prefix") + ":"
}
