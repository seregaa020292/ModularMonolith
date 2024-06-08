package repository

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pgsql"
)

type OwnerRepo struct {
	db *pgsql.DB
}

func NewOwnerRepo(db *pgsql.DB) *OwnerRepo {
	return &OwnerRepo{db: db}
}
