package app

import (
	"log/slog"
	"net"
	"net/http"

	"github.com/seregaa020292/ModularMonolith/internal/config/consts"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/router"
)

type App struct {
	router *router.Router
}

func New(router *router.Router) *App {
	return &App{
		router: router,
	}
}

func (app *App) Start() {
	serv := &http.Server{
		Addr:    net.JoinHostPort("", consts.ServerPort),
		Handler: app.router.Setup(),
	}

	slog.Info("starting http server on " + serv.Addr)
	if err := serv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (app *App) Stop() {

}
