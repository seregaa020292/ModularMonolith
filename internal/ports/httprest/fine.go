package httprest

import (
	"context"
	"time"

	"github.com/google/uuid"
	openapitypes "github.com/oapi-codegen/runtime/types"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
	"github.com/seregaa020292/ModularMonolith/pkg/utils/gog"
)

type FineHandler struct{}

func NewFineHandler() *FineHandler {
	return &FineHandler{}
}

func (h FineHandler) CreateFine(ctx context.Context, request openapi.CreateFineRequestObject) (openapi.CreateFineResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (h FineHandler) ListFines(ctx context.Context, request openapi.ListFinesRequestObject) (openapi.ListFinesResponseObject, error) {
	return openapi.ListFines200JSONResponse{
		{
			Id:          gog.Ptr(uuid.New()),
			VehicleId:   uuid.New(),
			Amount:      200,
			Description: gog.Ptr("Description"),
			Status:      "Status New",
			DueDate: openapitypes.Date{
				Time: time.Now(),
			},
			IssueDate: openapitypes.Date{
				Time: time.Now(),
			},
			CreatedAt: gog.Ptr(time.Now()),
			UpdatedAt: gog.Ptr(time.Now()),
		},
	}, nil
}
