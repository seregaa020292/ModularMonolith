package repository

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pg"
)

type ReportRepo struct {
	db *pg.DB
}

func NewReportRepo(db *pg.DB) *ReportRepo {
	return &ReportRepo{db: db}
}
