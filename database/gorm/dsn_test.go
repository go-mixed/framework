package gorm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	contractsdatabase "gopkg.in/go-mixed/framework/contracts/database"
	"gopkg.in/go-mixed/framework/testing/mock"
)

const (
	testConnection = "mysql"
	testHost       = "127.0.0.1"
	testPort       = 3306
	testDatabase   = "forge"
	testUsername   = "root"
	testPassword   = "123123"
)

var testConfig = contractsdatabase.Config{
	Host:     testHost,
	Port:     testPort,
	Database: testDatabase,
	Username: testUsername,
	Password: testPassword,
}

func TestMysqlDsn(t *testing.T) {
	charset := "utf8mb4"
	loc := "Local"
	mockConfig := mock.Config()
	mockConfig.On("GetString", fmt.Sprintf("database.connections.%s.charset", testConnection)).Return(charset).Once()
	mockConfig.On("GetString", fmt.Sprintf("database.connections.%s.loc", testConnection)).Return(loc).Once()

	assert.Equal(t, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		testUsername, testPassword, testHost, testPort, testDatabase, charset, true, loc), MysqlDsn(testConnection, testConfig))
}

func TestPostgresqlDsn(t *testing.T) {
	sslmode := "disable"
	timezone := "UTC"
	mockConfig := mock.Config()
	mockConfig.On("GetString", fmt.Sprintf("database.connections.%s.sslmode", testConnection)).Return(sslmode).Once()
	mockConfig.On("GetString", fmt.Sprintf("database.connections.%s.timezone", testConnection)).Return(timezone).Once()

	assert.Equal(t, fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		testHost, testUsername, testPassword, testDatabase, testPort, sslmode, timezone), PostgresqlDsn(testConnection, testConfig))
}

func TestSqliteDsn(t *testing.T) {
	assert.Equal(t, testDatabase, SqliteDsn(testConfig))
}

func TestSqlserverDsn(t *testing.T) {
	charset := "utf8mb4"
	mockConfig := mock.Config()
	mockConfig.On("GetString", fmt.Sprintf("database.connections.%s.charset", testConnection)).Return(charset).Once()

	assert.Equal(t, fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&charset=%s",
		testUsername, testPassword, testHost, testPort, testDatabase, charset), SqlserverDsn(testConnection, testConfig))
}
