package httprest

import (
	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/openapi"
)

var _ openapi.StrictServerInterface = (*OpenApiHandler)(nil)

var ModuleSet = wire.NewSet(
	NewFineHandler,
	NewNotificationHandler,
	NewOwnerHandler,
	NewPaymentHandler,
	NewVehicleHandler,
	wire.Struct(new(OpenApiHandler), "*"),

	NewAdminHandler,
	wire.Struct(new(AppApiHandler), "*"),
)

type OpenApiHandler struct {
	*FineHandler
	*NotificationHandler
	*OwnerHandler
	*PaymentHandler
	*VehicleHandler
}

type AppApiHandler struct {
	*AdminHandler
}
