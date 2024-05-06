package config

import (
	"fmt"
)

type ServerConfiguration struct {
	Port  string `mapstructure:"SERVER_PORT"`
	Host  string `mapstructure:"SERVER_HOST"`
	Debug bool   `mapstructure:"DEBUG"`
	// Secret string `mapsctructure:SECRET`" TODO: Add secret key
	Allowed_Hosts string `mapstructure:"ALLOWED_HOSTS"`

	// JWT Configuration for Auth0
	JWTAudience string `mapstructure:"AUTH0_AUDIENCE"`
	JWTDomain   string `mapstructure:"AUTH0_DOMAIN"`
}

func (s *ServerConfiguration) GenerateServerAddress() string {
	appServer := fmt.Sprintf("%s:%s", s.Host, s.Port)
	return appServer
}
