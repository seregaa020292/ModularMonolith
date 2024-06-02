package main

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/app"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/router"
	"github.com/seregaa020292/ModularMonolith/internal/ports/httprest"
)

func main() {
	r := router.NewRouter(httprest.New(
		httprest.NewFineHandler(),
		httprest.NewNotificationHandler(),
		httprest.NewOwnerHandler(),
		httprest.NewPaymentHandler(),
		httprest.NewVehicleHandler(),
	))
	newApp := app.New(r)
	newApp.Start()
	newApp.Stop()
}
