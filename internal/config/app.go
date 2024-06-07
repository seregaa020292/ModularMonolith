package config

import (
	"net"
	"strings"

	"github.com/seregaa020292/ModularMonolith/internal/config/consts"
)

type App struct {
	Name         string `env:"APP_NAME" env-default:""`
	Env          string `env:"ENV" env-default:"development"`
	CorsOrigins  string `env:"CORS_ORIGINS" env-default:""`
	LogFormatter string `env:"LOG_FORMATTER" env-default:"text"`
	LogLevel     string `env:"LOG_LEVEL" env-default:"info"`
}

func (a App) IsProduction() bool {
	return a.Env == "production"
}

func (a App) IsDevelopment() bool {
	return a.Env == "development"
}

func (a App) Addr() string {
	return net.JoinHostPort("", consts.ServerPort)
}

func (a App) AllowedOrigins() []string {
	return strings.Split(a.CorsOrigins, ",")
}
