package crypt

import (
	"gopkg.in/go-mixed/framework.v1/contracts/crypt"
)

type Crypt struct {
}

func NewCrypt() crypt.ICrypt {
	return NewAES()
}
