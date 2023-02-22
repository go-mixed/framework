package gorm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"gopkg.in/go-mixed/framework.v1/contracts/config/mocks"
	contractsdatabase "gopkg.in/go-mixed/framework.v1/contracts/database"
	"gopkg.in/go-mixed/framework.v1/contracts/database/orm"
	"gopkg.in/go-mixed/framework.v1/testing/mock"
)

func TestDialector(t *testing.T) {
	var mockConfig *mocks.Config
	host := "localhost"
	port := 3306
	database := "forge"
	username := "root"
	password := "123123"

	tests := []struct {
		description     string
		connection      orm.Driver
		setup           func()
		expectDialector gorm.Dialector
		expectErr       error
	}{
		{
			description: "mysql",
			connection:  orm.DriverMysql,
			setup: func() {
				mockConfig.On("GetString", "database.connections.mysql.driver").
					Return(orm.DriverMysql.String()).Once()
				mockConfig.On("GetString", "database.connections.mysql.charset").
					Return("utf8mb4").Once()
				mockConfig.On("GetString", "database.connections.mysql.loc").
					Return("Local").Once()
			},
			expectDialector: mysql.New(mysql.Config{
				DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
					username, password, host, port, database, "utf8mb4", true, "Local"),
			}),
		},
		{
			description: "postgresql",
			connection:  orm.DriverPostgresql,
			setup: func() {
				mockConfig.On("GetString", "database.connections.postgresql.driver").
					Return(orm.DriverPostgresql.String()).Once()
				mockConfig.On("GetString", "database.connections.postgresql.sslmode").
					Return("disable").Once()
				mockConfig.On("GetString", "database.connections.postgresql.timezone").
					Return("UTC").Once()
			},
			expectDialector: postgres.New(postgres.Config{
				DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
					host, username, password, database, port, "disable", "UTC"),
			}),
		},
		{
			description: "sqlite",
			connection:  orm.DriverSqlite,
			setup: func() {
				mockConfig.On("GetString", "database.connections.sqlite.driver").
					Return(orm.DriverSqlite.String()).Once()
			},
			expectDialector: sqlite.Open(database),
		},
		{
			description: "sqlserver",
			connection:  orm.DriverSqlserver,
			setup: func() {
				mockConfig.On("GetString", "database.connections.sqlserver.driver").
					Return(orm.DriverSqlserver.String()).Once()
				mockConfig.On("GetString", "database.connections.sqlserver.charset").
					Return("utf8mb4").Once()
			},
			expectDialector: sqlserver.New(sqlserver.Config{
				DSN: fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&charset=%s",
					username, password, host, port, database, "utf8mb4"),
			}),
		},
		{
			description: "error driver",
			connection:  "goravel",
			setup: func() {
				mockConfig.On("GetString", "database.connections.goravel.driver").
					Return("goravel").Once()
			},
			expectErr: errors.New(fmt.Sprintf("err database driver: %s, only support mysql, postgresql, sqlite and sqlserver", "goravel")),
		},
	}

	for _, test := range tests {
		mockConfig = mock.Config()
		test.setup()
		dialector, err := dialector(test.connection.String(), contractsdatabase.Config{
			Host:     host,
			Port:     port,
			Database: database,
			Username: username,
			Password: password,
		})
		assert.Equal(t, test.expectDialector, dialector)
		assert.Equal(t, test.expectErr, err)
	}
}
