package logger

import (
	"io"
	"log/slog"
	"os"

	"github.com/seregaa020292/ModularMonolith/internal/config"
)

type Logger struct {
	*slog.Logger
}

func New(cfg config.App) *Logger {
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

	return &Logger{Logger: logger}
}
