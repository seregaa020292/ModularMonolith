package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/seregaa020292/ModularMonolith/pkg/utils/gog"
)

type (
	RequestLogger struct {
		logger *slog.Logger
	}
	entryLogger struct {
		req    *http.Request
		logger *slog.Logger
	}
)

func GetEntryLogger(ctx context.Context) *slog.Logger {
	if entry, ok := ctx.Value(chimiddleware.LogEntryCtxKey).(*entryLogger); ok {
		return entry.logger
	}
	return slog.Default()
}

func SetEntryLoggerCtxFromWriter(w http.ResponseWriter) context.Context {
	return context.WithValue(context.Background(),
		chimiddleware.LogEntryCtxKey,
		slog.With(slog.String("correlation_id", GetCorrelationIDResponse(w))),
	)
}

func NewRequestLogger(logger *slog.Logger) *RequestLogger {
	return &RequestLogger{
		logger: logger,
	}
}

func (l RequestLogger) NewLogEntry(r *http.Request) chimiddleware.LogEntry {
	return &entryLogger{
		req: r,
		logger: l.logger.With(
			slog.String("request_id", chimiddleware.GetReqID(r.Context())),
			slog.String("correlation_id", GetCorrelationID(r.Context())),
		),
	}
}

func (l entryLogger) Write(status, bytes int, header http.Header, elapsed time.Duration, extra any) {
	log := l.logger.With(
		slog.String("url", fmt.Sprintf("%s://%s%s %s",
			gog.If(l.req.TLS != nil, "https", "http"), l.req.Host, l.req.RequestURI, l.req.Proto)),
		slog.String("method", l.req.Method),
		slog.String("user_agent", l.req.UserAgent()),
		slog.Int("status", status),
		slog.Int("bytes", bytes),
		slog.Float64("elapsed", float64(elapsed.Nanoseconds())/1000000.0),
	)

	switch {
	case status >= 500:
		log.Error("HTTP Server Error")
	case status >= 400:
		log.Warn("HTTP Client Error")
	default:
		//log.Info("HTTP Request Processed")
	}
}

func (l entryLogger) Panic(v any, stack []byte) {
	l.logger.Error("HTTP Handler Panic",
		slog.String("stack", string(stack)),
		slog.String("panic", fmt.Sprintf("%+v", v)),
	)
}
