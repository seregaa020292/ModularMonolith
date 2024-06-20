// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"context"
	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/internal/fine/query"
	"github.com/seregaa020292/ModularMonolith/internal/fine/repository"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/http/respond"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/http/router"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/logger"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/pgsql"
	repository2 "github.com/seregaa020292/ModularMonolith/internal/notification/repository"
	repository3 "github.com/seregaa020292/ModularMonolith/internal/owner/repository"
	repository4 "github.com/seregaa020292/ModularMonolith/internal/payment/repository"
	"github.com/seregaa020292/ModularMonolith/internal/ports/httprest"
	"log/slog"
)

// Injectors from service_provider.go:

// NewServiceProvider функция использует Google Wire для автоматической сборки зависимостей.
//
// В качестве параметров принимает контекст выполнения ctx и конфигурацию cfg.
// Возвращает указатель на serviceProvider, функцию для очистки и ошибку, если таковая возникнет.
func NewServiceProvider(ctx context.Context, cfg config.Config) (*serviceProvider, func(), error) {
	pg := cfg.PG
	db, cleanup, err := pgsql.New(pg)
	if err != nil {
		return nil, nil, err
	}
	fineRepo := repository.NewFineRepo(db)
	getListHandler := query.NewGetList(fineRepo)
	fineHandler := httprest.NewFineHandler(getListHandler)
	notificationRepo := repository2.NewNotificationRepo(db)
	notificationHandler := httprest.NewNotificationHandler(notificationRepo)
	ownerRepo := repository3.NewOwnerRepo(db)
	ownerHandler := httprest.NewOwnerHandler(ownerRepo)
	paymentRepo := repository4.NewPaymentRepo(db)
	paymentHandler := httprest.NewPaymentHandler(paymentRepo)
	vehicleHandler := httprest.NewVehicleHandler()
	openApiHandler := &httprest.OpenApiHandler{
		FineHandler:         fineHandler,
		NotificationHandler: notificationHandler,
		OwnerHandler:        ownerHandler,
		PaymentHandler:      paymentHandler,
		VehicleHandler:      vehicleHandler,
	}
	handle := respond.New()
	adminHandler := httprest.NewAdminHandler(handle)
	appApiHandler := &httprest.AppApiHandler{
		AdminHandler: adminHandler,
	}
	app := cfg.App
	slogLogger := logger.NewSlog(app)
	routerRouter, err := router.New(openApiHandler, appApiHandler, handle, slogLogger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	appServiceProvider := &serviceProvider{
		Router: routerRouter,
		Logger: slogLogger,
	}
	return appServiceProvider, func() {
		cleanup()
	}, nil
}

// service_provider.go:

type serviceProvider struct {
	Router *router.Router
	Logger *slog.Logger
}
