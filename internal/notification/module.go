package notification

import (
	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/notification/repository"
)

var Module = wire.NewSet(
	repository.NewNotificationRepo,
)
