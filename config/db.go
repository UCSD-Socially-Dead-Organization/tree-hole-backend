package config

import (
	"fmt"
)

type DatabaseConfiguration struct {
	Dbname   string `mapstructure:"DB_NAME"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	LogMode  bool   `mapstructure:"DB_LOG_MODE"`
	sslMode  string `mapstructure:"SSL_MODE"`
}

// GetDSN returns the Data Source Name for the database connection. Currently, it is using PostgreSQL, if we have more than one DB then the name should change.
func (d *DatabaseConfiguration) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		d.Host, d.User, d.Password, d.Dbname, d.Port, d.sslMode,
	)
}
