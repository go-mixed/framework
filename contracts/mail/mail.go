package mail

//go:generate mockery --name=Mail
type IMail interface {
	Content(content Content) IMail
	From(address From) IMail
	To(addresses []string) IMail
	Cc(addresses []string) IMail
	Bcc(addresses []string) IMail
	Attach(files []string) IMail
	Send() error
	Queue(queue *Queue) error
}

type Content struct {
	Subject string
	Html    string
}

type Queue struct {
	Connection string
	Queue      string
}

type From struct {
	Address string
	Name    string
}
