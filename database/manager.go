package database

import (
	"context"
	configfacade "gopkg.in/go-mixed/framework.v1/facades/config"
	"gopkg.in/go-mixed/framework.v1/support/manager"

	ormcontract "gopkg.in/go-mixed/framework.v1/contracts/database/orm"
)

type DatabaseManager struct {
	manager.Manager[ormcontract.IOrm]
}

func NewDatabaseManager() *DatabaseManager {
	m := &DatabaseManager{}
	m.Manager = manager.MakeManager[ormcontract.IOrm](m.DefaultDriverName, m.makeConnection)
	return m
}

func (m *DatabaseManager) Connection(name string) ormcontract.IOrm {
	return m.MustDriver(name)
}

func (m *DatabaseManager) DefaultDriverName() string {
	return configfacade.GetString("database.default")
}

func (m *DatabaseManager) makeConnection(connectionName string) (ormcontract.IOrm, error) {
	return NewDatabase(context.Background(), connectionName)
}
