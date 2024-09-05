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
	cfg    *config.Config
	closer *closer.Closer
}

func New(cfg *config.Config) *App {
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
	defer app.gracefulStop()

	sp, clean, err := NewServiceProvider(ctx, app.cfg)
	if err != nil {
		panic(err)
	}

	app.addCloser(clean)

	serv := &http.Server{
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		Addr:              app.cfg.App.Addr(),
		Handler:           sp.Router.Setup(app.cfg.App),
	}

	sp.Logger.Info("Starting server",
		slog.String("app", app.cfg.App.Name),
		slog.String("addr", serv.Addr),
	)
	if err := serv.ListenAndServe(); err != nil {
		sp.Logger.Error(err.Error())
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
