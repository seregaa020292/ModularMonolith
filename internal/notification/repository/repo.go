package repository

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pgsql"
)

type NotificationRepo struct {
	db *pgsql.DB
}

func NewNotificationRepo(db *pgsql.DB) *NotificationRepo {
	return &NotificationRepo{db: db}
}
