package mail

import (
	"context"
	configinstance "gopkg.in/go-mixed/framework.v1/config"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"log"
	"testing"
	"time"

	"github.com/gookit/color"
	"github.com/stretchr/testify/suite"

	"gopkg.in/go-mixed/framework.v1/contracts/event"
	"gopkg.in/go-mixed/framework.v1/contracts/mail"
	queuecontract "gopkg.in/go-mixed/framework.v1/contracts/queue"
	"gopkg.in/go-mixed/framework.v1/facades"
	"gopkg.in/go-mixed/framework.v1/queue"
	"gopkg.in/go-mixed/framework.v1/support/file"
	testingdocker "gopkg.in/go-mixed/framework.v1/testing/docker"
	"gopkg.in/go-mixed/framework.v1/testing/mock"
)

type ApplicationTestSuite struct {
	suite.Suite
}

func TestApplicationTestSuite(t *testing.T) {
	if !file.Exists("../.env") {
		color.Redln("No mail tests run, need create .env based on .env.example, then initialize it")
		return
	}

	redisPool, redisResource, err := testingdocker.Redis()
	if err != nil {
		log.Fatalf("Get redis error: %s", err)
	}

	initConfig(redisResource.GetPort("6379/tcp"))
	facades.Mail = NewApplication()
	suite.Run(t, new(ApplicationTestSuite))

	if err := redisPool.Purge(redisResource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func (s *ApplicationTestSuite) SetupTest() {

}

func (s *ApplicationTestSuite) TestSendMailBy25Port() {
	config.Add("mail", map[string]any{
		"host": config.Env("MAIL_HOST", ""),
		"port": 25,
		"from": map[string]any{
			"address": config.Env("MAIL_FROM_ADDRESS", "hello@example.com"),
			"name":    config.Env("MAIL_FROM_NAME", "Example"),
		},
		"username": config.Env("MAIL_USERNAME"),
		"password": config.Env("MAIL_PASSWORD"),
	})
	s.Nil(facades.Mail.To([]string{config.Env("MAIL_TO").(string)}).
		Cc([]string{config.Env("MAIL_CC").(string)}).
		Bcc([]string{config.Env("MAIL_BCC").(string)}).
		Attach([]string{"../logo.png"}).
		Content(mail.Content{Subject: "Goravel Test", Html: "<h1>Hello Goravel</h1>"}).
		Send())
}

func (s *ApplicationTestSuite) TestSendMailBy465Port() {
	config.Add("mail", map[string]any{
		"host": config.Env("MAIL_HOST", ""),
		"port": 465,
		"from": map[string]any{
			"address": config.Env("MAIL_FROM_ADDRESS", "hello@example.com"),
			"name":    config.Env("MAIL_FROM_NAME", "Example"),
		},
		"username": config.Env("MAIL_USERNAME"),
		"password": config.Env("MAIL_PASSWORD"),
	})
	s.Nil(facades.Mail.To([]string{config.Env("MAIL_TO").(string)}).
		Cc([]string{config.Env("MAIL_CC").(string)}).
		Bcc([]string{config.Env("MAIL_BCC").(string)}).
		Attach([]string{"../logo.png"}).
		Content(mail.Content{Subject: "Goravel Test", Html: "<h1>Hello Goravel</h1>"}).
		Send())
}

func (s *ApplicationTestSuite) TestSendMailBy587Port() {
	s.Nil(facades.Mail.To([]string{config.Env("MAIL_TO").(string)}).
		Cc([]string{config.Env("MAIL_CC").(string)}).
		Bcc([]string{config.Env("MAIL_BCC").(string)}).
		Attach([]string{"../logo.png"}).
		Content(mail.Content{Subject: "Goravel Test", Html: "<h1>Hello Goravel</h1>"}).
		Send())
}

func (s *ApplicationTestSuite) TestSendMailWithFrom() {
	s.Nil(facades.Mail.From(mail.From{Address: config.GetString("mail.from.address"), Name: config.GetString("mail.from.name")}).
		To([]string{config.Env("MAIL_TO").(string)}).
		Cc([]string{config.Env("MAIL_CC").(string)}).
		Bcc([]string{config.Env("MAIL_BCC").(string)}).
		Attach([]string{"../logo.png"}).
		Content(mail.Content{Subject: "Goravel Test With From", Html: "<h1>Hello Goravel</h1>"}).
		Send())
}

func (s *ApplicationTestSuite) TestQueueMail() {
	facades.Queue = queue.NewApplication()
	facades.Queue.Register([]queuecontract.Job{
		&SendMailJob{},
	})

	mockEvent, _ := mock.Event()
	mockEvent.On("GetEvents").Return(map[event.Event][]event.Listener{}).Once()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		s.Nil(facades.Queue.Worker(nil).Run())

		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	time.Sleep(3 * time.Second)
	s.Nil(facades.Mail.To([]string{config.Env("MAIL_TO").(string)}).
		Cc([]string{config.Env("MAIL_CC").(string)}).
		Bcc([]string{config.Env("MAIL_BCC").(string)}).
		Attach([]string{"../logo.png"}).
		Content(mail.Content{Subject: "Goravel Test Queue", Html: "<h1>Hello Goravel</h1>"}).
		Queue(nil))
	time.Sleep(1 * time.Second)

	mockEvent.AssertExpectations(s.T())
}

func initConfig(redisPort string) {
	newConfig := configinstance.NewConfig("../.env")
	newConfig.Add("app", map[string]any{
		"name": "goravel",
	})
	newConfig.Add("mail", map[string]any{
		"host": newConfig.Env("MAIL_HOST", ""),
		"port": newConfig.Env("MAIL_PORT", 587),
		"from": map[string]any{
			"address": newConfig.Env("MAIL_FROM_ADDRESS", "hello@example.com"),
			"name":    newConfig.Env("MAIL_FROM_NAME", "Example"),
		},
		"username": newConfig.Env("MAIL_USERNAME"),
		"password": newConfig.Env("MAIL_PASSWORD"),
	})
	newConfig.Add("queue", map[string]any{
		"default": "redis",
		"connections": map[string]any{
			"sync": map[string]any{
				"driver": "sync",
			},
			"redis": map[string]any{
				"driver":     "redis",
				"connection": "default",
				"queue":      "default",
			},
		},
	})
	newConfig.Add("database", map[string]any{
		"redis": map[string]any{
			"default": map[string]any{
				"host":     "localhost",
				"password": "",
				"port":     redisPort,
				"database": 0,
			},
		},
	})

	container.Instance("config", newConfig)
}
