package repository

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pg"
)

type OwnerRepo struct {
	db *pg.DB
}

func NewOwnerRepo(db *pg.DB) *OwnerRepo {
	return &OwnerRepo{db: db}
}
