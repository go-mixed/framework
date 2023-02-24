package db

import (
	"context"
	"database/sql"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/database/orm"
)

func getOrm() orm.IOrm {
	return container.MustMake[orm.IOrm]("db")
}

func Connection(name string) orm.IOrm {
	return getOrm().Connection(name)
}

func DB() (*sql.DB, error) {
	return getOrm().DB()
}

func Query() orm.DB {
	return getOrm().Query()
}

func Transaction(txFunc func(tx orm.Transaction) error) error {
	return getOrm().Transaction(txFunc)
}

func WithContext(ctx context.Context) orm.IOrm {
	return getOrm().WithContext(ctx)
}
