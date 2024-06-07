package response

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/errs"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/http/middleware"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/logger"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
)

type ErrorResponse struct {
	logger *logger.Logger
}

func NewErrorResponse(logger *logger.Logger) *ErrorResponse {
	return &ErrorResponse{
		logger: logger,
	}
}

func (e ErrorResponse) Send(ctx context.Context, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/problem+json")

	var errCustom errs.ErrorCustomer
	if errors.As(err, &errCustom) {
		e.logger.Error(errCustom.Error(),
			slog.String("request_id", chimiddleware.GetReqID(ctx)),
			slog.String("correlation_id", w.Header().Get(middleware.CorrelationIDHeader)),
			slog.String("stack", fmt.Sprintf("%+v", errCustom.OriginalError())),
		)

		w.WriteHeader(errCustom.StatusCode())
		if err := json.NewEncoder(w).Encode(openapi.Error{
			Code:    int32(errCustom.StatusCode()),
			Message: errCustom.Error(),
		}); err != nil {
			e.logger.Error(err.Error())
		}
		return
	}

	e.logger.Error(http.StatusText(http.StatusInternalServerError),
		slog.String("request_id", chimiddleware.GetReqID(ctx)),
		slog.String("correlation_id", w.Header().Get(middleware.CorrelationIDHeader)),
		slog.String("stack", fmt.Sprintf("%+v", err)),
	)

	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(openapi.Error{
		Code:    0,
		Message: http.StatusText(http.StatusInternalServerError),
	}); err != nil {
		e.logger.Error(err.Error())
	}
}
