package app

import (
	"context"
	"os"
	"syscall"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/internal/config/di"
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

func (a App) Run(ctx context.Context) {
	container, clean, err := di.New(ctx, a.cfg)
	if err != nil {
		panic(err)
	}

	defer a.closer.Wait()
	a.closer.Add(func() error {
		clean()
		return nil
	})

	container.Server.Run(ctx, a.cfg.App)
}
