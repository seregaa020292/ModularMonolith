package logger

import (
	"io"
	"log/slog"
	"os"

	log "github.com/seregaa020292/ModularMonolith/internal/infrastructure/logger"
	"github.com/seregaa020292/ModularMonolith/pkg/prettylog"
)

type Config struct {
	Formatter string `env:"LOG_FORMATTER" env-default:"text"`
	Level     string `env:"LOG_LEVEL" env-default:"info"`
}

func New(cfg Config) *slog.Logger {
	opts := log.SlogOptions(cfg.Level)
	writer := io.MultiWriter(os.Stdout)

	var handler slog.Handler
	switch cfg.Formatter {
	case "json":
		handler = slog.NewJSONHandler(writer, opts)
	case "pretty":
		handler = prettylog.New(writer, opts)
	default:
		handler = slog.NewTextHandler(writer, opts)
	}

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return logger
}
