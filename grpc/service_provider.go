package grpc

import (
	"gopkg.in/go-mixed/framework.v1/container"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Register() {
	container.Singleton((*Application)(nil), func(args ...any) (any, error) {
		return NewApplication(), nil
	})
	container.Alias("grpc", (*Application)(nil))
	//facades.Grpc = NewApplication()
}

func (route *ServiceProvider) Boot() {

}
