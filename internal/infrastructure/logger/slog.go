package logger

import (
	"io"
	"log/slog"
	"os"

	"github.com/pkg/errors"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/pkg/prettylog"
)

func NewSlog(cfg config.App) *slog.Logger {
	var level slog.Level
	if err := level.UnmarshalText([]byte(cfg.LogLevel)); err != nil {
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level:       level,
		ReplaceAttr: replaceAttr,
	}

	var handler slog.Handler
	writer := io.MultiWriter(os.Stdout)
	switch cfg.LogFormatter {
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

func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	switch a.Value.Kind() {
	case slog.KindAny:
		switch v := a.Value.Any().(type) {
		case error:
			a.Value = fmtErr(v)
		}
	}

	return a
}

func fmtErr(err error) slog.Value {
	groupValues := []slog.Attr{
		slog.String("msg", err.Error()),
	}

	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	var st stackTracer
	if ok := errors.As(err, &st); ok {
		groupValues = append(groupValues, slog.Any("trace", st.StackTrace()[:10]))
	}

	return slog.GroupValue(groupValues...)
}
