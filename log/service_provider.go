package log

import "gopkg.in/go-mixed/framework/facades"

type ServiceProvider struct {
}

func (log *ServiceProvider) Register() {
	facades.Log = NewLogrusApplication()
}

func (log *ServiceProvider) Boot() {

}
