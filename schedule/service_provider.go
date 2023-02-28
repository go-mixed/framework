package schedule

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/schedule"
)

type ServiceProvider struct {
}

func (sp *ServiceProvider) Register() {
	container.Singleton((schedule.ISchedule)(nil), func(args ...any) (any, error) {
		return NewApplication(), nil
	})
	container.Alias("schedule", (schedule.ISchedule)(nil))
}

func (sp *ServiceProvider) Boot() {

}
