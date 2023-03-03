package mail

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/mail"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
	queue2 "gopkg.in/go-mixed/framework.v1/facades/queue"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Register() {
	container.Singleton((mail.IMail)(nil), func(args ...any) (any, error) {
		return NewMail(), nil
	})
	container.Alias("mail", (mail.IMail)(nil))
}

func (route *ServiceProvider) Boot() {
	queue2.Register([]queue.IJob{
		&SendMailJob{},
	}...)
}
