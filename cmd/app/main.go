package main

import (
	"context"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/app"
)

func main() {
	app.New(config.MustNew()).
		Run(context.Background())
}
