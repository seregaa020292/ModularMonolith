package repository

import "database/sql"

type FineRepo struct {
	db *sql.DB
}

func NewFineRepo(db *sql.DB) *FineRepo {
	return &FineRepo{db: db}
}
