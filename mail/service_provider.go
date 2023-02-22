package mail

import (
	"gopkg.in/go-mixed/framework/contracts/queue"
	"gopkg.in/go-mixed/framework/facades"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Register() {
	facades.Mail = NewApplication()
}

func (route *ServiceProvider) Boot() {
	facades.Queue.Register([]queue.Job{
		&SendMailJob{},
	})
}
