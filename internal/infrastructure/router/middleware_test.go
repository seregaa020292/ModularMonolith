package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCorrelationID(t *testing.T) {
	const autoGenerated = "autoGenerated"

	tests := []struct {
		name    string
		headers map[string]string
		want    string
	}{
		{
			name: "Define Correlation-ID",
			headers: map[string]string{
				"Correlation-ID": "123",
			},
			want: "123",
		},
		{
			name: "Define X-Correlation-ID",
			headers: map[string]string{
				"X-Correlation-ID": "123",
			},
			want: "123",
		},
		{
			name: "Auto generated Correlation-ID",
			want: autoGenerated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/", nil)
			assert.NoError(t, err)

			for k, v := range tt.headers {
				req.Header.Set(k, v)
			}

			rr := httptest.NewRecorder()
			handler := NewCorrelationID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			handler.ServeHTTP(rr, req)

			correlationID := rr.Header().Get("Correlation-ID")

			if tt.want == autoGenerated {
				assert.NotZero(t, correlationID)
				return
			}

			assert.Equal(t, tt.want, correlationID)
		})
	}
}
