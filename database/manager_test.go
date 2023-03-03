package database

import (
	"context"
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"

	ormcontract "gopkg.in/go-mixed/framework.v1/contracts/database/orm"
	"gopkg.in/go-mixed/framework.v1/database/gorm"
	"gopkg.in/go-mixed/framework.v1/database/orm"
	"gopkg.in/go-mixed/framework.v1/support/file"
)

var connections = []ormcontract.Driver{
	ormcontract.DriverMysql,
	ormcontract.DriverPostgresql,
	ormcontract.DriverSqlite,
	ormcontract.DriverSqlserver,
}

type User struct {
	orm.Model
	orm.SoftDeletes
	Name   string
	Avatar string
}

type OrmSuite struct {
	suite.Suite
}

var (
	testMysqlDB      ormcontract.DB
	testPostgresqlDB ormcontract.DB
	testSqliteDB     ormcontract.DB
	testSqlserverDB  ormcontract.DB
)

func TestOrmSuite(t *testing.T) {
	mysqlPool, mysqlDocker, mysqlDB, err := gorm.MysqlDocker()
	if err != nil {
		log.Fatalf("Get mysql error: %s", err)
	}
	testMysqlDB = mysqlDB

	postgresqlPool, postgresqlDocker, postgresqlDB, err := gorm.PostgresqlDocker()
	if err != nil {
		log.Fatalf("Get postgresql error: %s", err)
	}
	testPostgresqlDB = postgresqlDB

	_, _, sqliteDB, err := gorm.SqliteDocker("laravel")
	if err != nil {
		log.Fatalf("Get sqlite error: %s", err)
	}
	testSqliteDB = sqliteDB

	sqlserverPool, sqlserverDocker, sqlserverDB, err := gorm.SqlserverDocker()
	if err != nil {
		log.Fatalf("Get sqlserver error: %s", err)
	}
	testSqlserverDB = sqlserverDB

	suite.Run(t, new(OrmSuite))

	file.Remove("laravel")

	if err := mysqlPool.Purge(mysqlDocker); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	if err := postgresqlPool.Purge(postgresqlDocker); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	if err := sqlserverPool.Purge(sqlserverDocker); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func (s *OrmSuite) SetupTest() {

}

func (s *OrmSuite) TestConnection() {
	testOrm := newTestManager()
	for _, connection := range connections {
		s.NotNil(testOrm.Connection(connection.String()))
	}
}

func (s *OrmSuite) TestDB() {
	testOrm := newTestManager()
	db, err := testOrm.MustDefaultDriver().DB()
	s.NotNil(db)
	s.Nil(err)

	for _, connection := range connections {
		db, err := testOrm.Connection(connection.String()).DB()
		s.NotNil(db)
		s.Nil(err)
	}
}

func (s *OrmSuite) TestQuery() {
	testOrm := newTestManager()
	s.NotNil(testOrm.MustDefaultDriver().Query())

	s.NotPanics(func() {
		for i := 0; i < 5; i++ {
			go func() {
				var user User
				_ = testOrm.MustDefaultDriver().Query().Find(&user, 1)
			}()
		}
	})

	for _, connection := range connections {
		s.NotNil(testOrm.Connection(connection.String()).Query())
	}
}

func (s *OrmSuite) TestTransactionSuccess() {
	testOrm := newTestManager()
	for _, connection := range connections {
		user := User{Name: "transaction_success_user", Avatar: "transaction_success_avatar"}
		user1 := User{Name: "transaction_success_user1", Avatar: "transaction_success_avatar1"}
		s.Nil(testOrm.Connection(connection.String()).Transaction(func(tx ormcontract.Transaction) error {
			s.Nil(tx.Create(&user))
			s.Nil(tx.Create(&user1))

			return nil
		}))

		var user2, user3 User
		s.Nil(testOrm.Connection(connection.String()).Query().Find(&user2, user.ID))
		s.Nil(testOrm.Connection(connection.String()).Query().Find(&user3, user1.ID))
	}
}

func (s *OrmSuite) TestTransactionError() {
	testOrm := newTestManager()
	for _, connection := range connections {
		s.NotNil(testOrm.Connection(connection.String()).Transaction(func(tx ormcontract.Transaction) error {
			user := User{Name: "transaction_error_user", Avatar: "transaction_error_avatar"}
			s.Nil(tx.Create(&user))

			user1 := User{Name: "transaction_error_user1", Avatar: "transaction_error_avatar1"}
			s.Nil(tx.Create(&user1))

			return errors.New("error")
		}))

		var users []User
		s.Nil(testOrm.Connection(connection.String()).Query().Find(&users))
		s.Equal(0, len(users))
	}
}

func newTestManager() *DatabaseManager {
	m := NewDatabaseManager()

	m.Extend(ormcontract.DriverMysql.String(), func(name string, args ...any) (ormcontract.IOrm, error) {
		return WrapDB(context.Background(), testMysqlDB), nil
	})

	m.Extend(ormcontract.DriverPostgresql.String(), func(name string, args ...any) (ormcontract.IOrm, error) {
		return WrapDB(context.Background(), testPostgresqlDB), nil
	})

	m.Extend(ormcontract.DriverSqlite.String(), func(name string, args ...any) (ormcontract.IOrm, error) {
		return WrapDB(context.Background(), testSqliteDB), nil
	})

	m.Extend(ormcontract.DriverSqlserver.String(), func(name string, args ...any) (ormcontract.IOrm, error) {
		return WrapDB(context.Background(), testSqlserverDB), nil
	})

	return m
}
