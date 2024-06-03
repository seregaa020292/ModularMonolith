package config

import (
	"net"

	"github.com/seregaa020292/ModularMonolith/internal/config/consts"
)

type App struct {
	Name string `env:"APP_NAME" env-default:""`
	Env  string `env:"ENV" env-default:"development"`
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
