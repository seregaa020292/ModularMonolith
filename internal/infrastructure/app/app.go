package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"syscall"
	"time"

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
	app.addCloser(clean)

	serv := &http.Server{
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		Addr:              app.cfg.App.Addr(),
		Handler:           provide.Router.Setup(app.cfg.App),
	}

	provide.Logger.Info("Starting server",
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
