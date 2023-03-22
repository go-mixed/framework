package mail

import (
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
)

type SendMailJob struct {
}

var _ queue.IJob = (*SendMailJob)(nil)

// Signature The name and signature of the job.
func (r *SendMailJob) Signature() string {
	return "laravel_send_mail_job"
}

// Handle Execute the job.
func (r *SendMailJob) Handle(args ...any) error {
	return SendMail(
		args[0].(string),
		args[1].(string),
		args[2].(string),
		args[3].(string),
		args[4].([]string),
		args[5].([]string),
		args[6].([]string),
		args[7].([]string),
	)
}
