package main

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/app"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/router"
	"github.com/seregaa020292/ModularMonolith/internal/ports/httprest"
)

func main() {
	newApp := app.New(router.NewRouter(httprest.HttpRest{}))
	newApp.Start()
	newApp.Stop()
}
