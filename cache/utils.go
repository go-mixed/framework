package cache

import (
	"gopkg.in/go-mixed/framework.v1/facades"
)

func prefix() string {
	return facades.Config.GetString("cache.prefix") + ":"
}
