package database

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"gopkg.in/go-mixed/framework.v1/container"
	ormcontract "gopkg.in/go-mixed/framework.v1/contracts/database/orm"
	databasegorm "gopkg.in/go-mixed/framework.v1/database/gorm"
)

type Orm struct {
	ctx      context.Context
	instance ormcontract.DB
}

var _ ormcontract.IOrm = (*Orm)(nil)

func NewOrm(ctx context.Context, connectionName string) (*Orm, error) {
	gormDB, err := databasegorm.NewDB(ctx, connectionName)
	if err != nil {
		return nil, err
	}

	return &Orm{ctx: ctx, instance: gormDB}, nil
}

func WrapDB(ctx context.Context, db ormcontract.DB) *Orm {
	return &Orm{ctx: ctx, instance: db}
}

func (r *Orm) Connection(connectionName string) ormcontract.IOrm {
	return container.MustMake[*ConnectionManager]("database.manager").Connection(connectionName).WithContext(r.ctx)
}

func (r *Orm) DB() (*sql.DB, error) {
	db := r.Query().(*databasegorm.DB)

	return db.Gorm().DB()
}

func (r *Orm) Query() ormcontract.DB {
	return r.instance
}

func (r *Orm) Transaction(txFunc func(tx ormcontract.Transaction) error) error {
	tx, err := r.Query().Begin()
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

func (r *Orm) WithContext(ctx context.Context) ormcontract.IOrm {
	db := r.Query().(*databasegorm.DB)

	return &Orm{
		ctx:      ctx,
		instance: db.WithContext(ctx),
	}
}
