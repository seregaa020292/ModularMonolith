package query

import (
	"context"

	"github.com/seregaa020292/ModularMonolith/internal/models/app/public/model"
	"github.com/seregaa020292/ModularMonolith/pkg/decorator"
)

type (
	Repository interface {
		List(ctx context.Context) ([]*model.Fines, error)
	}
	PayloadGetList struct{}
	GetListHandler decorator.QueryHandler[PayloadGetList, []*model.Fines]
)

type GetList struct {
	repo Repository
}

func NewGetList(repo Repository) GetListHandler {
	return &GetList{
		repo: repo,
	}
}

func (g GetList) Handle(ctx context.Context, payload PayloadGetList) ([]*model.Fines, error) {
	return g.repo.List(ctx)
}
