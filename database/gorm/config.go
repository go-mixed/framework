package gorm

import (
	"fmt"

	contractsdatabase "gopkg.in/go-mixed/framework/contracts/database"
	contractsorm "gopkg.in/go-mixed/framework/contracts/database/orm"
	"gopkg.in/go-mixed/framework/facades"
)

func Configs(connection string) (readConfigs, writeConfigs []contractsdatabase.Config, err error) {
	readConfigs = getReadConfigs(connection)
	writeConfigs = getWriteConfigs(connection)
	if len(readConfigs) == 0 && len(writeConfigs) == 0 {
		return nil, nil, nil
	}

	return
}

func getReadConfigs(connection string) []contractsdatabase.Config {
	configs := facades.Config.Get(fmt.Sprintf("database.connections.%s.read", connection))
	if c, exist := configs.([]contractsdatabase.Config); exist {
		return fillDefaultForConfigs(connection, c)
	}

	return []contractsdatabase.Config{}
}

func getWriteConfigs(connection string) []contractsdatabase.Config {
	configs := facades.Config.Get(fmt.Sprintf("database.connections.%s.write", connection))
	if c, exist := configs.([]contractsdatabase.Config); exist {
		return fillDefaultForConfigs(connection, c)
	}

	return fillDefaultForConfigs(connection, []contractsdatabase.Config{{}})
}

func fillDefaultForConfigs(connection string, configs []contractsdatabase.Config) []contractsdatabase.Config {
	var newConfigs []contractsdatabase.Config
	driver := facades.Config.GetString(fmt.Sprintf("database.connections.%s.driver", connection))
	for _, config := range configs {
		if driver != contractsorm.DriverSqlite.String() {
			if config.Host == "" {
				config.Host = facades.Config.GetString(fmt.Sprintf("database.connections.%s.host", connection))
			}
			if config.Port == 0 {
				config.Port = facades.Config.GetInt(fmt.Sprintf("database.connections.%s.port", connection))
			}
			if config.Username == "" {
				config.Username = facades.Config.GetString(fmt.Sprintf("database.connections.%s.username", connection))
			}
			if config.Password == "" {
				config.Password = facades.Config.GetString(fmt.Sprintf("database.connections.%s.password", connection))
			}
		}
		if config.Database == "" {
			config.Database = facades.Config.GetString(fmt.Sprintf("database.connections.%s.database", connection))
		}
		newConfigs = append(newConfigs, config)
	}

	return newConfigs
}
