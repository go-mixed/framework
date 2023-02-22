package route

import (
	"gopkg.in/go-mixed/framework.v1/facades"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Register() {
	facades.Route = NewGin()
}

func (route *ServiceProvider) Boot() {

}
