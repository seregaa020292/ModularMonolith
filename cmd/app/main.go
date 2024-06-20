package main

import (
	"context"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/app"
)

func main() {
	cfg := config.MustNew()

	a := app.New(cfg)
	a.Run(context.Background())
}
