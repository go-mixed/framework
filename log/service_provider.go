package log

import "gopkg.in/go-mixed/framework.v1/facades"

type ServiceProvider struct {
}

func (log *ServiceProvider) Register() {
	facades.Log = NewLogrusApplication()
}

func (log *ServiceProvider) Boot() {

}
