package httprest

import (
	"context"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
	"github.com/seregaa020292/ModularMonolith/internal/notification/repository"
)

type NotificationHandler struct{}

func NewNotificationHandler(repo *repository.NotificationRepo) *NotificationHandler {
	return &NotificationHandler{}
}

func (h NotificationHandler) CreateNotification(ctx context.Context, request openapi.CreateNotificationRequestObject) (openapi.CreateNotificationResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (h NotificationHandler) ListNotifications(ctx context.Context, request openapi.ListNotificationsRequestObject) (openapi.ListNotificationsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
