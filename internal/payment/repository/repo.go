package repository

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pgsql"
)

type PaymentRepo struct {
	db *pgsql.DB
}

func NewPaymentRepo(db *pgsql.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}
