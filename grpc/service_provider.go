package grpc

import (
	"gopkg.in/go-mixed/framework.v1/facades"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Register() {
	facades.Grpc = NewApplication()
}

func (route *ServiceProvider) Boot() {

}
