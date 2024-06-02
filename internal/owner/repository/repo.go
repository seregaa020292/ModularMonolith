package repository

import "database/sql"

type OwnerRepo struct {
	db *sql.DB
}

func NewOwnerRepo(db *sql.DB) *OwnerRepo {
	return &OwnerRepo{db: db}
}
