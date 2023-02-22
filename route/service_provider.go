package route

import (
	"gopkg.in/go-mixed/framework/facades"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Register() {
	facades.Route = NewGin()
}

func (route *ServiceProvider) Boot() {

}
