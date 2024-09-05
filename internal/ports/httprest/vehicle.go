package httprest

import (
	"context"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/openapi"
)

type VehicleHandler struct{}

func NewVehicleHandler() *VehicleHandler {
	return &VehicleHandler{}
}

func (h VehicleHandler) CreateVehicle(ctx context.Context, request openapi.CreateVehicleRequestObject) (openapi.CreateVehicleResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (h VehicleHandler) ListVehicles(ctx context.Context, request openapi.ListVehiclesRequestObject) (openapi.ListVehiclesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
