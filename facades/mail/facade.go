package mail

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/mail"
)

func getMail() mail.IMail {
	return container.MustMake[mail.IMail]("mail")
}

func Content(content mail.Content) mail.IMail {
	return getMail().Content(content)
}

func From(address mail.From) mail.IMail {
	return getMail().From(address)
}

func To(addresses []string) mail.IMail {
	return getMail().To(addresses)
}

func Cc(addresses []string) mail.IMail {
	return getMail().Cc(addresses)
}

func Bcc(addresses []string) mail.IMail {
	return getMail().Bcc(addresses)
}

func Attach(files []string) mail.IMail {
	return getMail().Attach(files)
}

func Send() error {
	return getMail().Send()
}

func Queue(queue *mail.Queue) error {
	return getMail().Queue(queue)
}
