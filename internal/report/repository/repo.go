package repository

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pgsql"
)

type ReportRepo struct {
	db *pgsql.DB
}

func NewReportRepo(db *pgsql.DB) *ReportRepo {
	return &ReportRepo{db: db}
}
