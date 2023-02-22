package filesystem

import (
	"gopkg.in/go-mixed/framework/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Storage = NewStorage()
}

func (database *ServiceProvider) Boot() {

}
