package repository

import (
	"context"
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pg"
	"github.com/seregaa020292/ModularMonolith/internal/models/app/public/model"
	"github.com/seregaa020292/ModularMonolith/pkg/utils/gog"
)

type FineRepo struct {
	db *pg.DB
}

func NewFineRepo(db *pg.DB) *FineRepo {
	return &FineRepo{db: db}
}

func (r FineRepo) List(ctx context.Context) ([]*model.Fines, error) {
	return []*model.Fines{
		{
			ID:          uuid.New(),
			VehicleID:   gog.Ptr(uuid.New()),
			IssueDate:   time.Now(),
			DueDate:     time.Now(),
			Amount:      200,
			Status:      "Status New",
			Description: gog.Ptr("Description"),
			CreatedAt:   time.Now(),
			UpdatedAt:   gog.Ptr(time.Now()),
		},
	}, gog.If(rand.IntN(2) == 0, errors.New("some error"), nil)
}
