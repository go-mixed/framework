package log

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"gopkg.in/go-mixed/framework.v1/support/manager"
)

type LogManager struct {
	manager.Manager[*Logger]
}

func NewChannelManager() *LogManager {
	m := &LogManager{}
	m.Manager = manager.MakeManager[*Logger](m.DefaultDriverName, m.makeChannel)
	return m
}

func (m *LogManager) DefaultDriverName() string {
	return config.GetString("logging.default")
}

func (m *LogManager) Channel(channelName string) *Logger {
	return m.MustDriver(channelName)
}

func (m *LogManager) makeChannel(channelName string) (*Logger, error) {
	return NewLogger(context.Background(), channelName)
}
