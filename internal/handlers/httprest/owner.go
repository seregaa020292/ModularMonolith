package httprest

import (
	"context"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/openapi"
	"github.com/seregaa020292/ModularMonolith/internal/owner/repository"
)

type OwnerHandler struct{}

func NewOwnerHandler(repo *repository.OwnerRepo) *OwnerHandler {
	return &OwnerHandler{}
}

func (h OwnerHandler) CreateOwner(ctx context.Context, request openapi.CreateOwnerRequestObject) (openapi.CreateOwnerResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (h OwnerHandler) ListOwners(ctx context.Context, request openapi.ListOwnersRequestObject) (openapi.ListOwnersResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
