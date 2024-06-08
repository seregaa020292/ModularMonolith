package logger

import (
	"io"
	"log/slog"
	"os"

	"github.com/seregaa020292/ModularMonolith/internal/config"
)

func New(cfg config.App) *slog.Logger {
	var (
		level   slog.Level
		handler slog.Handler
		writer  = io.MultiWriter(os.Stdout)
	)

	if err := level.UnmarshalText([]byte(cfg.LogLevel)); err != nil {
		level = slog.LevelInfo
	}

	switch cfg.LogFormatter {
	case "json":
		handler = slog.NewJSONHandler(writer, &slog.HandlerOptions{
			Level: level,
		})
	default:
		handler = slog.NewTextHandler(writer, &slog.HandlerOptions{
			Level: level,
		})
	}

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return logger
}
