package main

import (
	"context"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/app"
	"github.com/seregaa020292/ModularMonolith/pkg/utils/gog"
)

func main() {
	cfg := gog.Must(config.New())

	newApp := app.New(cfg)
	newApp.Run(context.Background())
}
