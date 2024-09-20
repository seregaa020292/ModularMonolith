package httprest

import (
	"context"

	"github.com/seregaa020292/ModularMonolith/internal/fine/query"
	"github.com/seregaa020292/ModularMonolith/internal/handlers/httprest/presenter"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/errs"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/openapi"
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
	//if err := h.validator.Validate(request.Body); err != nil {
	//	return nil, errs.NewValidationError("Неверный запрос", err)
	//}

	return openapi.CreateFine201Response{}, nil
}

func (h FineHandler) ListFines(ctx context.Context, request openapi.ListFinesRequestObject) (openapi.ListFinesResponseObject, error) {
	fines, err := h.getList.Handle(ctx, query.PayloadGetList(request))
	if err != nil {
		return nil, errs.NewDomainError("Не удалось получить список штрафов", err)
	}

	return presenter.ListFines(fines), nil
}
