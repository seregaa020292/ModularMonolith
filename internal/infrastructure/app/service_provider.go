//go:build wireinject
// +build wireinject

package app

import (
	"context"

	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/internal/fine"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/http/response"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/http/router"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/logger"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pgsql"
	"github.com/seregaa020292/ModularMonolith/internal/notification"
	"github.com/seregaa020292/ModularMonolith/internal/owner"
	"github.com/seregaa020292/ModularMonolith/internal/payment"
	"github.com/seregaa020292/ModularMonolith/internal/ports/httprest"
)

type serviceProvider struct {
	Router *router.Router
	Logger *logger.Logger
}

// NewServiceProvider функция использует Google Wire для автоматической сборки зависимостей.
//
// В качестве параметров принимает контекст выполнения ctx и конфигурацию cfg.
// Возвращает указатель на serviceProvider, функцию для очистки и ошибку, если таковая возникнет.
func NewServiceProvider(ctx context.Context, cfg config.Config) (*serviceProvider, func(), error) {
	panic(wire.Build(
		wire.FieldsOf(new(config.Config), "App", "PG"),

		logger.New,
		pgsql.New,
		response.NewErrorResponse,

		fine.ModuleSet,
		notification.ModuleSet,
		owner.ModuleSet,
		payment.ModuleSet,

		httprest.NewFineHandler,
		httprest.NewNotificationHandler,
		httprest.NewOwnerHandler,
		httprest.NewPaymentHandler,
		httprest.NewVehicleHandler,
		httprest.NewAdminHandler,
		wire.Struct(new(httprest.HttpRest), "*"),
		router.New,

		wire.Struct(new(serviceProvider), "*"),
	))
}
