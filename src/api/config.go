package api

import (
	"github.com/spf13/viper"
)

// Config api config
type Config struct {
	// The port to bind the web application server to
	Port int
}

// InitConfig initial config
func InitConfig() (*Config, error) {
	config := &Config{
		Port: viper.GetInt("port"),
	}

	if config.Port == 0 {
		config.Port = 3030
	}

	return config, nil
}
