package schedule

import (
	"github.com/robfig/cron/v3"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
	"gopkg.in/go-mixed/framework.v1/facades/log"

	"gopkg.in/go-mixed/framework.v1/contracts/schedule"
	"gopkg.in/go-mixed/framework.v1/schedule/support"
)

type Schedule struct {
	cron *cron.Cron
}

func NewSchedule() *Schedule {
	return &Schedule{}
}

func (app *Schedule) Call(callback func()) schedule.Event {
	return &support.Event{Callback: callback}
}

func (app *Schedule) Command(command string) schedule.Event {
	return &support.Event{Command: command}
}

func (app *Schedule) Register(events []schedule.Event) {
	if app.cron == nil {
		app.cron = cron.New(cron.WithChain(cron.Recover(log.CronLog())), cron.WithLogger(log.CronLog()))
	}

	app.addEvents(events)
}

func (app *Schedule) Run() {
	app.cron.Start()
}

func (app *Schedule) addEvents(events []schedule.Event) {
	for _, event := range events {
		chain := cron.NewChain(cron.Recover(log.CronLog()))
		if event.GetDelayIfStillRunning() {
			chain = cron.NewChain(cron.Recover(log.CronLog()), cron.DelayIfStillRunning(log.CronLog()))
		} else if event.GetSkipIfStillRunning() {
			chain = cron.NewChain(cron.Recover(log.CronLog()), cron.SkipIfStillRunning(log.CronLog()))
		}
		_, err := app.cron.AddJob(event.GetCron(), chain.Then(app.getJob(event)))

		if err != nil {
			log.Errorf("add schedule error: %v", err)
		}
	}
}

func (app *Schedule) getJob(event schedule.Event) cron.Job {
	return cron.FuncJob(func() {
		if event.GetCommand() != "" {
			artisan.Call(event.GetCommand())
		} else {
			event.GetCallback()()
		}
	})
}

//type Logger struct{}
//
//func (l *Logger) Info(msg string, keysAndValues ...any) {
//	color.Green.Printf("%s %v\n", msg, keysAndValues)
//}
//
//func (l *Logger) Error(err error, msg string, keysAndValues ...any) {
//	log.Error(msg, keysAndValues)
//}
