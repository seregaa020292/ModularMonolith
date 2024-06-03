package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App App
	PG  PG
}

func New() (Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		text, _ := cleanenv.GetDescription(cfg, nil)
		slog.Info(text)
		return Config{}, err
	}
	return cfg, nil
}
