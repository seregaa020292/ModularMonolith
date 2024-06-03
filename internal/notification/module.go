package notification

import (
	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/notification/repository"
)

var ModuleSet = wire.NewSet(
	repository.NewNotificationRepo,
)
