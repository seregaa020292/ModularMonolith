package httprest

import (
	"context"
	"time"

	"github.com/google/uuid"
	openapitypes "github.com/oapi-codegen/runtime/types"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
)

func (h HttpRest) ListFines(ctx context.Context, request openapi.ListFinesRequestObject) (openapi.ListFinesResponseObject, error) {
	return openapi.ListFines200JSONResponse{
		{
			Amount:      200,
			CreatedAt:   nil,
			Description: nil,
			DueDate: openapitypes.Date{
				Time: time.Now(),
			},
			Id: nil,
			IssueDate: openapitypes.Date{
				Time: time.Now(),
			},
			Status:    "",
			UpdatedAt: nil,
			VehicleId: uuid.New(),
		},
	}, nil
}
