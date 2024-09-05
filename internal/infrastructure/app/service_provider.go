//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"log/slog"

	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/internal/config/logger"
	"github.com/seregaa020292/ModularMonolith/internal/config/pg"
	"github.com/seregaa020292/ModularMonolith/internal/fine"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/respond"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/router"
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
func NewServiceProvider(ctx context.Context, cfg *config.Config) (*serviceProvider, func(), error) {
	panic(wire.Build(
		// Получения конфигурационных настроек
		wire.FieldsOf(new(*config.Config), "PG", "Logger"),

		// Инициализация компонентов
		pg.New,
		logger.New,
		respond.New,

		// Модули доменной логики
		fine.ModuleSet,
		notification.ModuleSet,
		owner.ModuleSet,
		payment.ModuleSet,

		// HTTP-обработчики
		httprest.ModuleSet,
		router.New,

		// Агрегатор всех сервисов и компонентов
		wire.Struct(new(serviceProvider), "*"),
	))
}
