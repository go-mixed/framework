package database

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"gopkg.in/go-mixed/framework.v1/container"
	ormcontract "gopkg.in/go-mixed/framework.v1/contracts/database/orm"
	databasegorm "gopkg.in/go-mixed/framework.v1/database/gorm"
)

type Database struct {
	ctx      context.Context
	instance ormcontract.DB
}

var _ ormcontract.IOrm = (*Database)(nil)

func NewDatabase(ctx context.Context, connectionName string) (*Database, error) {
	gormDB, err := databasegorm.NewDB(ctx, connectionName)
	if err != nil {
		return nil, err
	}

	return &Database{ctx: ctx, instance: gormDB}, nil
}

func WrapDB(ctx context.Context, db ormcontract.DB) *Database {
	return &Database{ctx: ctx, instance: db}
}

func (d *Database) Connection(connectionName string) ormcontract.IOrm {
	return container.MustMake[*ConnectionManager]("database.manager").Connection(connectionName).WithContext(d.ctx)
}

func (d *Database) DB() (*sql.DB, error) {
	db := d.Query().(*databasegorm.DB)

	return db.Gorm().DB()
}

func (d *Database) Query() ormcontract.DB {
	return d.instance
}

func (d *Database) Transaction(txFunc func(tx ormcontract.Transaction) error) error {
	tx, err := d.Query().Begin()
	if err != nil {
		return err
	}

	if err := txFunc(tx); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrapf(err, "rollback error: %v", err)
		}

		return err
	} else {
		return tx.Commit()
	}
}

func (d *Database) WithContext(ctx context.Context) ormcontract.IOrm {
	db := d.Query().(*databasegorm.DB)

	return &Database{
		ctx:      ctx,
		instance: db.WithContext(ctx),
	}
}
