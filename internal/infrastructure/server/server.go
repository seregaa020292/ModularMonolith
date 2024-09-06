package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/seregaa020292/ModularMonolith/internal/config/app"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/router"
)

type Server struct {
	router *router.Router
	logger *slog.Logger
}

func New(router *router.Router, logger *slog.Logger) *Server {
	return &Server{
		router: router,
		logger: logger,
	}
}

func (s *Server) Run(ctx context.Context, cfg app.Config) {
	serv := &http.Server{
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		Addr:              cfg.Addr(),
		Handler:           s.router.Setup(cfg),
	}

	s.logger.Info("Starting server",
		slog.String("app", cfg.Name),
		slog.String("addr", serv.Addr),
	)
	if err := serv.ListenAndServe(); err != nil {
		s.logger.Error(err.Error())
	}
}
