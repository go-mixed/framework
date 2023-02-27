package mail

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/mail"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
	"gopkg.in/go-mixed/framework.v1/facades"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Register() {
	container.Singleton((mail.IMail)(nil), func(args ...any) (any, error) {
		return NewApplication(), nil
	})
	container.Alias("mail", (mail.IMail)(nil))
	//facades.Mail = NewApplication()
}

func (route *ServiceProvider) Boot() {
	facades.Queue.Register([]queue.Job{
		&SendMailJob{},
	})
}
