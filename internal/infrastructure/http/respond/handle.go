package respond

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/pkg/errors"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/errs"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/http/middleware"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
	"github.com/seregaa020292/ModularMonolith/pkg/utils/gog"
)

type Handle struct{}

func New() *Handle {
	return &Handle{}
}

func (h Handle) Success(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h Handle) Error(ctx context.Context, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/problem+json")

	logger := middleware.GetEntryLogger(ctx)

	var errCustom errs.ErrorCustomer
	if errors.As(err, &errCustom) {
		lvlLogger := gog.If(errCustom.StatusCode() >= 500, logger.Error, logger.Warn)
		lvlLogger(errCustom.Error(), slog.Any("error", errCustom.OriginalError()))

		w.WriteHeader(errCustom.StatusCode())
		if err := json.NewEncoder(w).Encode(openapi.Error{
			Code:    int32(errCustom.StatusCode()),
			Message: errCustom.Error(),
		}); err != nil {
			logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	logger.Error(err.Error(), slog.Any("error", err))

	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(openapi.Error{
		Code:    0,
		Message: http.StatusText(http.StatusInternalServerError),
	}); err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
