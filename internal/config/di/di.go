//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/internal/config/logger"
	"github.com/seregaa020292/ModularMonolith/internal/config/pg"
	"github.com/seregaa020292/ModularMonolith/internal/fine"
	"github.com/seregaa020292/ModularMonolith/internal/handlers/httprest"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/respond"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/router"
	"github.com/seregaa020292/ModularMonolith/internal/notification"
	"github.com/seregaa020292/ModularMonolith/internal/owner"
	"github.com/seregaa020292/ModularMonolith/internal/payment"
)

type container struct {
	Server *server.Server
}

// NewRegistry функция использует Google Wire для автоматической сборки зависимостей.
//
// В качестве параметров принимает контекст выполнения ctx и конфигурацию cfg.
// Возвращает указатель на Registry, функцию для очистки и ошибку, если таковая возникнет.
func New(ctx context.Context, cfg *config.Config) (*container, func(), error) {
	panic(wire.Build(
		// Получения конфигурационных настроек
		wire.FieldsOf(new(*config.Config), "PG", "Logger"),

		// Инициализация компонентов
		pg.New,
		logger.New,
		respond.New,

		// Модули доменной логики
		fine.Module,
		notification.Module,
		owner.Module,
		payment.Module,

		// HTTP-обработчики
		httprest.Module,

		router.New,
		server.New,

		// Агрегатор всех сервисов и компонентов
		wire.Struct(new(container), "*"),
	))
}
