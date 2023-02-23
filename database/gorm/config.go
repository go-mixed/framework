package gorm

import (
	"fmt"
	contractsdatabase "gopkg.in/go-mixed/framework.v1/contracts/database"
	"gopkg.in/go-mixed/framework.v1/facades/config"

	contractsorm "gopkg.in/go-mixed/framework.v1/contracts/database/orm"
	configfacade "gopkg.in/go-mixed/framework.v1/facades/config"
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
	configs := config.Get(fmt.Sprintf("database.connections.%s.read", connection))
	if c, exist := configs.([]contractsdatabase.Config); exist {
		return fillDefaultForConfigs(connection, c)
	}

	return []contractsdatabase.Config{}
}

func getWriteConfigs(connection string) []contractsdatabase.Config {
	configs := config.Get(fmt.Sprintf("database.connections.%s.write", connection))
	if c, exist := configs.([]contractsdatabase.Config); exist {
		return fillDefaultForConfigs(connection, c)
	}

	return fillDefaultForConfigs(connection, []contractsdatabase.Config{{}})
}

func fillDefaultForConfigs(connection string, configs []contractsdatabase.Config) []contractsdatabase.Config {
	var newConfigs []contractsdatabase.Config
	driver := config.GetString(fmt.Sprintf("database.connections.%s.driver", connection))
	for _, config := range configs {
		if driver != contractsorm.DriverSqlite.String() {
			if config.Host == "" {
				config.Host = configfacade.GetString(fmt.Sprintf("database.connections.%s.host", connection))
			}
			if config.Port == 0 {
				config.Port = configfacade.GetInt(fmt.Sprintf("database.connections.%s.port", connection))
			}
			if config.Username == "" {
				config.Username = configfacade.GetString(fmt.Sprintf("database.connections.%s.username", connection))
			}
			if config.Password == "" {
				config.Password = configfacade.GetString(fmt.Sprintf("database.connections.%s.password", connection))
			}
		}
		if config.Database == "" {
			config.Database = configfacade.GetString(fmt.Sprintf("database.connections.%s.database", connection))
		}
		newConfigs = append(newConfigs, config)
	}

	return newConfigs
}
