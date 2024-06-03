package repository

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pg"
)

type FineRepo struct {
	db *pg.DB
}

func NewFineRepo(db *pg.DB) *FineRepo {
	return &FineRepo{db: db}
}
