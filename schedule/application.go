package schedule

import (
	"github.com/gookit/color"
	"github.com/robfig/cron/v3"
	"gopkg.in/go-mixed/framework.v1/facades/log"

	"gopkg.in/go-mixed/framework.v1/contracts/schedule"
	"gopkg.in/go-mixed/framework.v1/facades"
	"gopkg.in/go-mixed/framework.v1/schedule/support"
)

type Application struct {
	cron *cron.Cron
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Call(callback func()) schedule.Event {
	return &support.Event{Callback: callback}
}

func (app *Application) Command(command string) schedule.Event {
	return &support.Event{Command: command}
}

func (app *Application) Register(events []schedule.Event) {
	if app.cron == nil {
		app.cron = cron.New(cron.WithLogger(&Logger{}))
	}

	app.addEvents(events)
}

func (app *Application) Run() {
	app.cron.Start()
}

func (app *Application) addEvents(events []schedule.Event) {
	for _, event := range events {
		chain := cron.NewChain()
		if event.GetDelayIfStillRunning() {
			chain = cron.NewChain(cron.DelayIfStillRunning(&Logger{}))
		} else if event.GetSkipIfStillRunning() {
			chain = cron.NewChain(cron.SkipIfStillRunning(&Logger{}))
		}
		_, err := app.cron.AddJob(event.GetCron(), chain.Then(app.getJob(event)))

		if err != nil {
			log.Errorf("add schedule error: %v", err)
		}
	}
}

func (app *Application) getJob(event schedule.Event) cron.Job {
	return cron.FuncJob(func() {
		if event.GetCommand() != "" {
			facades.Artisan.Call(event.GetCommand())
		} else {
			event.GetCallback()()
		}
	})
}

type Logger struct{}

func (l *Logger) Info(msg string, keysAndValues ...any) {
	color.Green.Printf("%s %v\n", msg, keysAndValues)
}

func (l *Logger) Error(err error, msg string, keysAndValues ...any) {
	log.Error(msg, keysAndValues)
}
