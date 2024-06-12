package app

import (
	"goapp/common/postgresql"
)

type ConfigurationManagger struct {
	PostreSqlConfig postgresql.Config
}

func NewConfigurationManagger() *ConfigurationManagger {
	postgreSqlConfig := getPostreSqlConfig()
	return &ConfigurationManagger{
		PostreSqlConfig: postgreSqlConfig,
	}
}

func getPostreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "goapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
}
