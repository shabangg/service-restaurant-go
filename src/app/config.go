package app

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config app cconfig
type Config struct {
	SecretKey string
}

// InitConfig initial config
func InitConfig() (*Config, error) {
	config := &Config{
		SecretKey: viper.GetString("secretkey"),
	}

	if len(config.SecretKey) == 0 {
		return nil, fmt.Errorf("SecretKey must be set")
	}

	return config, nil
}