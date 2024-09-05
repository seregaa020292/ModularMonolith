package app

import (
	"net"
	"strings"

	"github.com/seregaa020292/ModularMonolith/internal/config/consts"
)

type Config struct {
	Name        string `env:"APP_NAME" env-default:""`
	Env         string `env:"ENV" env-default:"development"`
	CorsOrigins string `env:"CORS_ORIGINS" env-default:""`
}

func (c Config) IsProduction() bool {
	return c.Env == "production"
}

func (c Config) IsDevelopment() bool {
	return c.Env == "development"
}

func (c Config) Addr() string {
	return net.JoinHostPort("", consts.ServerPort)
}

func (c Config) AllowedOrigins() []string {
	return strings.Split(c.CorsOrigins, ",")
}
