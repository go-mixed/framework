package route

import (
	"gopkg.in/go-mixed/framework.v1/container"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Register() {
	container.Singleton((*Gin)(nil), func(args ...any) (any, error) {
		return NewGin(), nil
	})
	container.Alias("route", (*Gin)(nil))
	//facades.Route = NewGin()
}

func (route *ServiceProvider) Boot() {

}
