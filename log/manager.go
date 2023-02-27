package log

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"gopkg.in/go-mixed/framework.v1/support/manager"
)

type ChannelManager struct {
	manager.Manager[*Logger]
}

func NewChannelManager() *ChannelManager {
	m := &ChannelManager{}
	m.Manager = manager.MakeManager[*Logger](m.DefaultDriverName)
	return m
}

func (m *ChannelManager) DefaultDriverName() string {
	return config.GetString("logging.default")
}

func (m *ChannelManager) Channel(channelName string) *Logger {
	return m.MustDriver(channelName)
}

func (m *ChannelManager) extendChannels() {
	for name := range config.GetMap("logging.channels") {
		m.Extend(name, m.makeDriver)
	}
}

func (m *ChannelManager) makeDriver(chanelName string) (*Logger, error) {
	return NewLogger(context.Background(), chanelName)
}
