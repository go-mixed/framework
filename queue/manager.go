package queue

import (
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"gopkg.in/go-mixed/framework.v1/support/manager"
	"strings"
)

type QueueManager struct {
	manager.Manager[queue.IBroker]
}

func NewQueueManager() *QueueManager {
	m := &QueueManager{}
	m.Manager = manager.MakeManager[queue.IBroker](m.DefaultConnectionName, m.makeConnection)
	return m
}

func (m *QueueManager) Connection(connectionName string) queue.IBroker {
	return m.MustDriver(connectionName)
}

func (m *QueueManager) DefaultConnectionName() string {
	return config.GetString("queue.default")
}

func (m *QueueManager) makeConnection(connectionQueueName string) (queue.IBroker, error) {
	segments := strings.SplitN(connectionQueueName, "|", 2)
	var queueName string
	connectionName := segments[0]
	if connectionName == "" {
		connectionName = m.DefaultConnectionName()
	}
	if len(segments) == 1 {
		queueName = ""
	} else {
		queueName = segments[1]
	}

	driver := config.GetString("queue.connections."+connectionName+".driver", "")

	if m.HasCustomCreator(driver) {
		instance, err := m.CallCustomCreator(driver, connectionName, queueName)
		if err != nil {
			color.Redf("[Queue] Initialize queue driver \"%s.%s\" error: %v\n", connectionName, driver, err)
			return nil, errors.Errorf("[Cache] Initialize queue \"%s.%s\" error: %v\n", connectionName, driver, err)
		}

		return instance.(queue.IBroker), nil
	}

	color.Redf("[Queue] queue driver \"%s.%s\" is not defined.\n", connectionName, driver)
	return nil, errors.Errorf("[Queue] queue driver \"%s.%s\" is not defined.\n", connectionName, driver)
}
