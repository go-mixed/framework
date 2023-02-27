package mail

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"net/smtp"

	"github.com/jordan-wright/email"

	"gopkg.in/go-mixed/framework.v1/contracts/mail"
	queuecontract "gopkg.in/go-mixed/framework.v1/contracts/queue"
	"gopkg.in/go-mixed/framework.v1/facades"
)

type Application struct {
	clone    int
	content  mail.Content
	from     mail.From
	to       []string
	cc       []string
	bcc      []string
	attaches []string
}

func NewApplication() *Application {
	return &Application{}
}

func (r *Application) Content(content mail.Content) mail.IMail {
	instance := r.instance()
	instance.content = content

	return instance
}

func (r *Application) From(from mail.From) mail.IMail {
	instance := r.instance()
	instance.from = from

	return instance
}

func (r *Application) To(to []string) mail.IMail {
	instance := r.instance()
	instance.to = to

	return instance
}

func (r *Application) Cc(cc []string) mail.IMail {
	instance := r.instance()
	instance.cc = cc

	return instance
}

func (r *Application) Bcc(bcc []string) mail.IMail {
	instance := r.instance()
	instance.bcc = bcc

	return instance
}

func (r *Application) Attach(files []string) mail.IMail {
	instance := r.instance()
	instance.attaches = files

	return instance
}

func (r *Application) Send() error {
	return SendMail(r.content.Subject, r.content.Html, r.from.Address, r.from.Name, r.to, r.cc, r.bcc, r.attaches)
}

func (r *Application) Queue(queue *mail.Queue) error {
	job := facades.Queue.Job(&SendMailJob{}, []queuecontract.Arg{
		{Value: r.content.Subject, Type: "string"},
		{Value: r.content.Html, Type: "string"},
		{Value: r.from.Address, Type: "string"},
		{Value: r.from.Name, Type: "string"},
		{Value: r.to, Type: "[]string"},
		{Value: r.cc, Type: "[]string"},
		{Value: r.bcc, Type: "[]string"},
		{Value: r.attaches, Type: "[]string"},
	})
	if queue != nil {
		if queue.Connection != "" {
			job.OnConnection(queue.Connection)
		}
		if queue.Queue != "" {
			job.OnQueue(queue.Queue)
		}
	}

	return job.Dispatch()
}

func (r *Application) instance() *Application {
	if r.clone == 0 {
		return &Application{clone: 1}
	}

	return r
}

func SendMail(subject, html string, fromAddress, fromName string, to, cc, bcc, attaches []string) error {
	e := email.NewEmail()
	if fromAddress == "" {
		e.From = fmt.Sprintf("%s <%s>", config.GetString("mail.from.name"), config.GetString("mail.from.address"))
	} else {
		e.From = fmt.Sprintf("%s <%s>", fromName, fromAddress)
	}

	e.To = to
	e.Bcc = bcc
	e.Cc = cc
	e.Subject = subject
	e.HTML = []byte(html)

	for _, attach := range attaches {
		if _, err := e.AttachFile(attach); err != nil {
			return err
		}
	}

	port := config.GetInt("mail.port")
	switch port {
	case 465:
		return e.SendWithTLS(fmt.Sprintf("%s:%s", config.GetString("mail.host"), config.GetString("mail.port")),
			LoginAuth(config.GetString("mail.username"), config.GetString("mail.password")),
			&tls.Config{ServerName: config.GetString("mail.host")})
	case 587:
		return e.SendWithStartTLS(fmt.Sprintf("%s:%s", config.GetString("mail.host"), config.GetString("mail.port")),
			LoginAuth(config.GetString("mail.username"), config.GetString("mail.password")),
			&tls.Config{ServerName: config.GetString("mail.host")})
	default:
		return e.Send(fmt.Sprintf("%s:%d", config.GetString("mail.host"), port),
			LoginAuth(config.GetString("mail.username"), config.GetString("mail.password")))
	}
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}
