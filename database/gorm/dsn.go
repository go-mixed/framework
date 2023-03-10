package gorm

import (
	"fmt"

	contractsdatabase "gopkg.in/go-mixed/framework.v1/contracts/database"
	configfacade "gopkg.in/go-mixed/framework.v1/facades/config"
)

func MysqlDsn(connection string, config contractsdatabase.Config) string {
	host := config.Host
	if host == "" {
		return ""
	}

	charset := configfacade.GetString("database.connections." + connection + ".charset")
	loc := configfacade.GetString("database.connections." + connection + ".loc")

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		config.Username, config.Password, host, config.Port, config.Database, charset, true, loc)
}

func PostgresqlDsn(connection string, config contractsdatabase.Config) string {
	host := config.Host
	if host == "" {
		return ""
	}

	sslmode := configfacade.GetString("database.connections." + connection + ".sslmode")
	timezone := configfacade.GetString("database.connections." + connection + ".timezone")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		host, config.Username, config.Password, config.Database, config.Port, sslmode, timezone)
}

func SqliteDsn(config contractsdatabase.Config) string {
	return config.Database
}

func SqlserverDsn(connection string, config contractsdatabase.Config) string {
	host := config.Host
	if host == "" {
		return ""
	}

	charset := configfacade.GetString("database.connections." + connection + ".charset")

	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&charset=%s",
		config.Username, config.Password, host, config.Port, config.Database, charset)
}
