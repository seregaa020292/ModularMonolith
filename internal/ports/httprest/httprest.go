package httprest

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
)

var _ openapi.StrictServerInterface = (*OpenapiHandler)(nil)

type OpenapiHandler struct {
	*FineHandler
	*NotificationHandler
	*OwnerHandler
	*PaymentHandler
	*VehicleHandler
}

type AppHandler struct {
	*AdminHandler
}

type ServerHandler struct {
	Openapi OpenapiHandler
	App     AppHandler
}
