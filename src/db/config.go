package db

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config databse config
type Config struct {
	DatabaseURI string
	Database    string
}

func InitConfig() (error, *Config) {
	config := &Config{
		DatabaseURI: viper.GetString("databaseuri"),
		Database:    viper.GetString("database"),
	}
	if config.DatabaseURI == "" {
		return fmt.Errorf("DatabaseURI must be set"), nil
	}

	return nil, config
}
