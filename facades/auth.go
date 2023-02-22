package facades

import (
	"gopkg.in/go-mixed/framework.v1/contracts/auth"
	"gopkg.in/go-mixed/framework.v1/contracts/auth/access"
)

var (
	Auth auth.Auth
	Gate access.Gate
)
