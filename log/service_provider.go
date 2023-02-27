package log

import (
	"gopkg.in/go-mixed/framework.v1/container"
	contractslog "gopkg.in/go-mixed/framework.v1/contracts/log"
)

type ServiceProvider struct {
}

func (log *ServiceProvider) Register() {
	container.Singleton((*ChannelManager)(nil), func(args ...any) (any, error) {
		manager := NewChannelManager()
		manager.extendChannels()
		return manager, nil
	})
	container.Alias("log.manager", (*ChannelManager)(nil))

	container.Singleton((*contractslog.ILog)(nil), func(args ...any) (any, error) {
		return container.MustMake[*ChannelManager]("log.manager").MustDefaultDriver(), nil
	})
	container.Alias(contractslog.ILog(nil), (*contractslog.ILog)(nil))
	container.Alias("log", (*contractslog.ILog)(nil))
}

func (log *ServiceProvider) Boot() {

}
