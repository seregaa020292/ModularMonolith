package presenter

import (
	openapitypes "github.com/oapi-codegen/runtime/types"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/openapi"
	"github.com/seregaa020292/ModularMonolith/internal/models/app/public/model"
	"github.com/seregaa020292/ModularMonolith/pkg/gog"
)

func ListFines(fines []*model.Fines) openapi.ListFines200JSONResponse {
	response := make(openapi.ListFines200JSONResponse, len(fines))

	for i, fine := range fines {
		response[i] = openapi.Fine{
			Id:          gog.Ptr(fine.ID),
			VehicleId:   *fine.VehicleID,
			Amount:      float32(fine.Amount),
			Description: fine.Description,
			Status:      fine.Status,
			DueDate:     openapitypes.Date{Time: fine.DueDate},
			IssueDate:   openapitypes.Date{Time: fine.IssueDate},
			CreatedAt:   gog.Ptr(fine.CreatedAt),
			UpdatedAt:   fine.UpdatedAt,
		}
	}

	return response
}
