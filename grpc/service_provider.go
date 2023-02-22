package grpc

import (
	"gopkg.in/go-mixed/framework/facades"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Register() {
	facades.Grpc = NewApplication()
}

func (route *ServiceProvider) Boot() {

}
