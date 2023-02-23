package database

import (
	"context"
	"database/sql"
	"fmt"
	configfacade "gopkg.in/go-mixed/framework.v1/facades/config"

	"github.com/gookit/color"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	ormcontract "gopkg.in/go-mixed/framework.v1/contracts/database/orm"
	databasegorm "gopkg.in/go-mixed/framework.v1/database/gorm"
)

type Orm struct {
	ctx       context.Context
	instance  ormcontract.DB
	instances map[string]ormcontract.DB
}

func NewOrm(ctx context.Context) *Orm {
	defaultConnection := configfacade.GetString("database.default")
	gormDB, err := databasegorm.NewDB(ctx, defaultConnection)
	if err != nil {
		color.Redln(fmt.Sprintf("[Orm] Initialize %s connection error: %v", defaultConnection, err))

		return nil
	}
	if gormDB == nil {
		return nil
	}

	return &Orm{
		ctx:      ctx,
		instance: gormDB,
		instances: map[string]ormcontract.DB{
			defaultConnection: gormDB,
		},
	}
}

// DEPRECATED: use gorm.New()
func NewGormInstance(connection string) (*gorm.DB, error) {
	return databasegorm.New(connection)
}

func (r *Orm) Connection(name string) ormcontract.Orm {
	if name == "" {
		name = configfacade.GetString("database.default")
	}
	if instance, exist := r.instances[name]; exist {
		return &Orm{
			ctx:       r.ctx,
			instance:  instance,
			instances: r.instances,
		}
	}

	gormDB, err := databasegorm.NewDB(r.ctx, name)
	if err != nil || gormDB == nil {
		color.Redln(fmt.Sprintf("[Orm] Initialize %s connection error: %v", name, err))

		return nil
	}

	r.instances[name] = gormDB

	return &Orm{
		ctx:       r.ctx,
		instance:  gormDB,
		instances: r.instances,
	}
}

func (r *Orm) DB() (*sql.DB, error) {
	db := r.Query().(*databasegorm.DB)

	return db.Instance().DB()
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

func (r *Orm) WithContext(ctx context.Context) ormcontract.Orm {
	return NewOrm(ctx)
}
