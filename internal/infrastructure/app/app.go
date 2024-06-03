package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"syscall"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/pkg/closer"
)

type App struct {
	cfg    config.Config
	closer *closer.Closer
}

func New(cfg config.Config) *App {
	return &App{
		cfg: cfg,
		closer: closer.New(
			os.Interrupt,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		),
	}
}

func (app App) Run(ctx context.Context) {
	provide, clean, err := NewServiceProvider(ctx, app.cfg)
	if err != nil {
		panic(err)
	}

	defer app.gracefulStop()

	app.closer.Add(func() error {
		clean()
		return nil
	})

	serv := &http.Server{
		Addr:    app.cfg.App.Addr(),
		Handler: provide.Router.Setup(),
	}

	slog.Info(fmt.Sprintf(`starting app "%s" on %s`, app.cfg.App.Name, serv.Addr))
	if err := serv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (app App) gracefulStop() {
	app.closer.CloseAll()
	app.closer.Wait()
}
