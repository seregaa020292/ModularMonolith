package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	CorrelationIDHeader  = "Correlation-ID"
	XCorrelationIDHeader = "X-Correlation-ID"
)

type correlationIDKey struct{}

func CorrelationID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(CorrelationIDHeader)
		if id == "" {
			id = r.Header.Get(XCorrelationIDHeader)
		}
		if id == "" {
			id = uuid.NewString()
		}

		ctx := context.WithValue(r.Context(), correlationIDKey{}, id)

		w.Header().Set(CorrelationIDHeader, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetCorrelationID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if id, ok := ctx.Value(correlationIDKey{}).(string); ok {
		return id
	}
	return ""
}

func GetCorrelationIDResponse(w http.ResponseWriter) string {
	return w.Header().Get(CorrelationIDHeader)
}
