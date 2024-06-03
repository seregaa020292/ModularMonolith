package repository

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pg"
)

type NotificationRepo struct {
	db *pg.DB
}

func NewNotificationRepo(db *pg.DB) *NotificationRepo {
	return &NotificationRepo{db: db}
}
