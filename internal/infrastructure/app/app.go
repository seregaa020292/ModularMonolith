package app

import (
	"context"
	"os"
	"syscall"

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

func (a App) Run(ctx context.Context) {
	registry, clean, err := NewRegistry(ctx, a.cfg)
	if err != nil {
		panic(err)
	}

	defer a.gracefulStop()
	a.addCloser(clean)

	registry.server.Run(ctx, a.cfg.App)
}

func (a App) addCloser(fn func()) {
	a.closer.Add(func() error {
		fn()
		return nil
	})
}

func (a App) gracefulStop() {
	a.closer.CloseAll()
	a.closer.Wait()
}
