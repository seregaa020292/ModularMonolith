//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"log/slog"

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
	Logger *slog.Logger
}

// NewServiceProvider функция использует Google Wire для автоматической сборки зависимостей.
//
// В качестве параметров принимает контекст выполнения ctx и конфигурацию cfg.
// Возвращает указатель на serviceProvider, функцию для очистки и ошибку, если таковая возникнет.
func NewServiceProvider(ctx context.Context, cfg config.Config) (*serviceProvider, func(), error) {
	panic(wire.Build(
		// Получения конфигурационных настроек
		wire.FieldsOf(new(config.Config), "App", "PG"),

		// Инициализация компонентов
		pgsql.New,
		logger.NewSlog,
		response.NewErrorHandle,

		// Модули доменной логики
		fine.ModuleSet,
		notification.ModuleSet,
		owner.ModuleSet,
		payment.ModuleSet,

		// HTTP-обработчики
		httprest.NewFineHandler,
		httprest.NewNotificationHandler,
		httprest.NewOwnerHandler,
		httprest.NewPaymentHandler,
		httprest.NewVehicleHandler,
		httprest.NewAdminHandler,
		wire.Struct(new(httprest.OpenapiHandler), "*"),
		wire.Struct(new(httprest.AppHandler), "*"),
		wire.Struct(new(httprest.ServerHandler), "*"),
		router.New,

		// Агрегатор всех сервисов и компонентов
		wire.Struct(new(serviceProvider), "*"),
	))
}
