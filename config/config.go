package config

import (
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/logger"
	"github.com/spf13/viper"
)

type Configuration struct {
	// with squash, the fields of the embedded struct are treated
	// as if they were part of the outer struct for the purposes of mapping keys to fields.
	Server ServerConfiguration   `mapstructure:",squash"`
	DB     DatabaseConfiguration `mapstructure:",squash"`
}

func SetupConfig() (*Configuration, error) {
	var configuration *Configuration

	// this will read the .env file and setup ALL the environment variables, we no longer need to read the environment variables manually
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error to reading config file, %s", err)
		return nil, err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Errorf("error to decode, %v", err)
		return nil, err
	}

	return configuration, nil
}
