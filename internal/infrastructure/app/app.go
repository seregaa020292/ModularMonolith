package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"syscall"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/pkg/closer"
	"github.com/seregaa020292/ModularMonolith/pkg/utils/gog"
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
	app.addCloser(clean)

	serv := &http.Server{
		Addr:    app.cfg.App.Addr(),
		Handler: gog.Must(provide.Router.Setup(ctx, app.cfg.App)),
	}

	provide.Logger.Info("starting server",
		slog.String("app", app.cfg.App.Name),
		slog.String("addr", serv.Addr),
	)
	if err := serv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (app App) addCloser(fn func()) {
	app.closer.Add(func() error {
		fn()
		return nil
	})
}

func (app App) gracefulStop() {
	app.closer.CloseAll()
	app.closer.Wait()
}
