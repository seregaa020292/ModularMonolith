package httprest

import (
	"context"

	"github.com/seregaa020292/ModularMonolith/internal/fine/query"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
	"github.com/seregaa020292/ModularMonolith/internal/ports/httprest/presenter"
)

type FineHandler struct {
	getList query.GetListHandler
}

func NewFineHandler(getList query.GetListHandler) *FineHandler {
	return &FineHandler{
		getList: getList,
	}
}

func (h FineHandler) CreateFine(ctx context.Context, request openapi.CreateFineRequestObject) (openapi.CreateFineResponseObject, error) {
	return openapi.CreateFine201Response{}, nil
}

func (h FineHandler) ListFines(ctx context.Context, request openapi.ListFinesRequestObject) (openapi.ListFinesResponseObject, error) {
	fines, err := h.getList.Handle(ctx, query.PayloadGetList(request))
	if err != nil {
		return presenter.ListFines400(openapi.Error{
			Code:    112,
			Message: err.Error(),
		}), nil
	}

	return presenter.ListFines(fines), nil
}
