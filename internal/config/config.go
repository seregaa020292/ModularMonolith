package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/seregaa020292/ModularMonolith/internal/config/app"
	"github.com/seregaa020292/ModularMonolith/internal/config/logger"
	"github.com/seregaa020292/ModularMonolith/internal/config/pg"
)

type Config struct {
	App    app.Config
	PG     pg.Config
	Logger logger.Config
}

func New() (*Config, error) {
	cfg := new(Config)
	if err := cleanenv.ReadEnv(cfg); err != nil {
		text, _ := cleanenv.GetDescription(cfg, nil)
		slog.Info(text)
		return nil, err
	}
	return cfg, nil
}

func MustNew() *Config {
	cfg, err := New()
	if err != nil {
		panic(err)
	}
	return cfg
}
