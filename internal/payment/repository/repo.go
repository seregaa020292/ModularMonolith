package repository

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pg"
)

type PaymentRepo struct {
	db *pg.DB
}

func NewPaymentRepo(db *pg.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}
