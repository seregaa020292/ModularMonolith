package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/seregaa020292/ModularMonolith/internal/config/consts"
	"github.com/seregaa020292/ModularMonolith/pkg/utils/gog"
	"github.com/seregaa020292/ModularMonolith/pkg/utils/sensitive"
)

type (
	RequestLogger struct {
		logger *slog.Logger
	}
	entryLogger struct {
		*RequestLogger
		logger *slog.Logger
		req    *http.Request
		body   map[string]any
	}
)

func NewRequestLogger(logger *slog.Logger) *RequestLogger {
	return &RequestLogger{
		logger: logger,
	}
}

func GetEntryLogger(ctx context.Context) *slog.Logger {
	if entry, ok := ctx.Value(chimiddleware.LogEntryCtxKey).(*entryLogger); ok {
		return entry.logger
	}
	return slog.Default()
}

func WrapCtxLogger(w http.ResponseWriter) context.Context {
	return context.WithValue(context.Background(),
		chimiddleware.LogEntryCtxKey,
		&entryLogger{
			logger: slog.With(slog.String("correlation_id", GetCorrelationIDResponse(w))),
		},
	)
}

func (l *RequestLogger) NewLogEntry(r *http.Request) chimiddleware.LogEntry {
	return &entryLogger{
		RequestLogger: l,
		logger: l.logger.With(
			slog.String("request_id", chimiddleware.GetReqID(r.Context())),
			slog.String("correlation_id", GetCorrelationID(r.Context())),
		),
		req:  r,
		body: sanitizeRequestBody(r, consts.SensitiveDataMask, consts.SensitiveFilerKeys),
	}
}

func (l *entryLogger) Write(status, bytes int, header http.Header, elapsed time.Duration, extra any) {
	log := l.logger.With(
		slog.String("url", fmt.Sprintf("%s://%s%s %s",
			gog.If(l.req.TLS != nil, "https", "http"), l.req.Host, l.req.RequestURI, l.req.Proto)),
		slog.String("method", l.req.Method),
		slog.Int("status", status),
		slog.Any("body", l.body),
		slog.String("user_agent", l.req.UserAgent()),
		slog.String("remote_addr", l.req.RemoteAddr),
		slog.String("referer", l.req.Referer()),
		slog.Float64("elapsed", float64(elapsed.Nanoseconds())/1000000.0),
		slog.Int("bytes", bytes),
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

func (l *entryLogger) Panic(v any, stack []byte) {
	l.logger.Error("HTTP Handler Panic",
		slog.String("stack", string(stack)),
		slog.String("panic", fmt.Sprintf("%+v", v)),
	)
}

// sanitizeRequestBody читает и санитизирует тело запроса, удаляя конфиденциальные данные.
// Возвращает карту с данными.
func sanitizeRequestBody(r *http.Request, mask string, filerKeys []string) map[string]any {
	if r.Body == nil {
		return nil
	}

	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	body, err := sensitive.MapUnmarshal(bodyBytes, mask, filerKeys)
	if err != nil {
		return nil
	}

	return body
}
