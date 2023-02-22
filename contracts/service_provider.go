package contracts

import "gopkg.in/go-mixed/framework.v1/contracts/container"

type IServiceProvider interface {
	//Boot any application services after register.
	Boot(container container.IContainer)
	//Register any application services.
	Register(container container.IContainer)
}
