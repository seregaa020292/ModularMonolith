//go:build wireinject
// +build wireinject

package app

import (
	"context"

	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/internal/fine/repository"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pg"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/router"
	"github.com/seregaa020292/ModularMonolith/internal/ports/httprest"
)

type serviceProvider struct {
	Router *router.Router
}

func NewServiceProvider(ctx context.Context, cfg config.Config) (*serviceProvider, func(), error) {
	panic(wire.Build(
		wire.FieldsOf(new(config.Config), "PG"),
		pg.New,
		repository.NewFineRepo,
		httprest.NewFineHandler,
		httprest.NewNotificationHandler,
		httprest.NewOwnerHandler,
		httprest.NewPaymentHandler,
		httprest.NewVehicleHandler,
		httprest.NewAdminHandler,
		httprest.New,
		router.NewRouter,
		wire.Struct(new(serviceProvider), "*"),
	))
}
