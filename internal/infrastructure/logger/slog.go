package logger

import (
	"log/slog"

	"github.com/pkg/errors"
)

func SlogOptions(lvl string) *slog.HandlerOptions {
	var level slog.Level
	if err := level.UnmarshalText([]byte(lvl)); err != nil {
		level = slog.LevelInfo
	}

	return &slog.HandlerOptions{
		Level:       level,
		ReplaceAttr: replaceAttr,
	}
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
