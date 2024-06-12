package httprest

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
)

var _ openapi.StrictServerInterface = (*ServerHandler)(nil)

type ServerHandler struct {
	*FineHandler
	*NotificationHandler
	*OwnerHandler
	*PaymentHandler
	*VehicleHandler

	*AdminHandler
}
