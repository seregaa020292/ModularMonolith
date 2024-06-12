package middleware

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/seregaa020292/ModularMonolith/internal/config/consts"
)

func Test_sanitizeRequestBody(t *testing.T) {
	type args struct {
		r         *http.Request
		mask      string
		filerKeys []string
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "should sanitize request body",
			args: args{
				r: &http.Request{
					Body: io.NopCloser(strings.NewReader(`{"password": "123"}`)),
				},
				mask:      consts.SensitiveDataMask,
				filerKeys: consts.SensitiveFilerKeys,
			},
			want: map[string]any{
				"password": consts.SensitiveDataMask,
			},
		},
		{
			name: "should not sanitize request body",
			args: args{
				r: &http.Request{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sanitizeRequestBody(tt.args.r, tt.args.mask, tt.args.filerKeys)
			assert.Equal(t, tt.want, got)
		})
	}
}
