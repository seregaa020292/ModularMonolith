package httprest

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
)

var _ openapi.StrictServerInterface = (*HttpRest)(nil)

type HttpRest struct {
	*FineHandler
	*NotificationHandler
	*OwnerHandler
	*PaymentHandler
	*VehicleHandler
}

func New(
	fineHandler *FineHandler,
	notificationHandler *NotificationHandler,
	ownerHandler *OwnerHandler,
	paymentHandler *PaymentHandler,
	vehicleHandler *VehicleHandler,
) *HttpRest {
	return &HttpRest{
		FineHandler:         fineHandler,
		NotificationHandler: notificationHandler,
		OwnerHandler:        ownerHandler,
		PaymentHandler:      paymentHandler,
		VehicleHandler:      vehicleHandler,
	}
}
