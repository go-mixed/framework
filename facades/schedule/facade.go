package schedule

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/schedule"
)

func getSchedule() schedule.ISchedule {
	return container.MustMake[schedule.ISchedule]("schedule")
}

func Call(callback func()) schedule.Event {
	return getSchedule().Call(callback)
}

func Command(command string) schedule.Event {
	return getSchedule().Command(command)
}

func Register(events []schedule.Event) {
	getSchedule().Register(events)
}

func Run() {
	getSchedule().Run()
}
