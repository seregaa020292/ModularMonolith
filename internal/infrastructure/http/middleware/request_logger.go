package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/seregaa020292/ModularMonolith/pkg/utils/gog"
)

type (
	RequestLogger struct{}
	entryLogger   struct {
		req   *http.Request
		attrs []any
	}
)

func NewRequestLogger() *RequestLogger {
	return &RequestLogger{}
}

func (l RequestLogger) NewLogEntry(r *http.Request) chimiddleware.LogEntry {
	return &entryLogger{
		req: r,
		attrs: []any{
			slog.String("request_id", chimiddleware.GetReqID(r.Context())),
			slog.String("correlation_id", GetCorrelationID(r.Context())),
			slog.String("scheme", gog.If(r.TLS != nil, "https", "http")),
			slog.String("host", r.Host),
			slog.String("ip", r.RemoteAddr),
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.String("user_agent", r.UserAgent()),
		},
	}
}

func (l entryLogger) Write(status, bytes int, header http.Header, elapsed time.Duration, extra any) {
	attrs := append(l.attrs,
		slog.Int("status", status),
		slog.Int("bytes", bytes),
		slog.Float64("elapsed", float64(elapsed.Nanoseconds())/1000000.0),
	)

	switch {
	case status >= 500:
		slog.Error("HTTP Server Error", attrs...)
	case status >= 400:
		slog.Warn("HTTP Client Error", attrs...)
	default:
		//slog.Info("HTTP Request Processed", attrs...)
	}
}

func (l entryLogger) Panic(v any, stack []byte) {
	slog.Error("HTTP Handler Panic",
		slog.String("request_id", chimiddleware.GetReqID(l.req.Context())),
		slog.String("correlation_id", GetCorrelationID(l.req.Context())),
		slog.String("stack", string(stack)),
		slog.String("panic", fmt.Sprintf("%+v", v)),
	)
}
