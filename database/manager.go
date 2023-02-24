package database

import (
	"context"
	configfacade "gopkg.in/go-mixed/framework.v1/facades/config"
	"gopkg.in/go-mixed/framework.v1/support/manager"

	ormcontract "gopkg.in/go-mixed/framework.v1/contracts/database/orm"
)

type ConnectionManager struct {
	manager.Manager[ormcontract.IOrm]
}

func NewConnectionManager() *ConnectionManager {
	orm := &ConnectionManager{}
	orm.Manager = manager.MakeManager[ormcontract.IOrm](orm.DefaultDriverName)
	return orm
}

func (m *ConnectionManager) Connection(name string) ormcontract.IOrm {
	return m.MustDriver(name)
}

func (m *ConnectionManager) DefaultDriverName() string {
	return configfacade.GetString("database.default")
}

func (m *ConnectionManager) makeDriver(connectionName string) (ormcontract.IOrm, error) {
	return NewOrm(context.Background(), connectionName)
}

func (m *ConnectionManager) extendConnections() {
	for name := range configfacade.GetMap("database.connections") {
		m.Extend(name, m.makeDriver)
	}
}
