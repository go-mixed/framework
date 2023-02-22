package schedule

import (
	"gopkg.in/go-mixed/framework/facades"
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register() {
	facades.Schedule = NewApplication()
}

func (receiver *ServiceProvider) Boot() {

}
