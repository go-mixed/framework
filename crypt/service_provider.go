package crypt

import (
	"gopkg.in/go-mixed/framework.v1/container"
	contractCrypt "gopkg.in/go-mixed/framework.v1/contracts/crypt"
)

type ServiceProvider struct {
}

func (sp *ServiceProvider) Register() {
	container.Singleton((*Crypt)(nil), func(args ...any) (any, error) {
		return NewCrypt(), nil
	})
	container.Alias("crypt", (*Crypt)(nil))
	container.Alias(contractCrypt.ICrypt(nil), (*Crypt)(nil))
}

func (sp *ServiceProvider) Boot() {

}
