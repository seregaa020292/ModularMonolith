package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/seregaa020292/ModularMonolith/internal/config/app"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/router"
)

type Option func(*http.Server)

type Server struct {
	cfg    app.Config
	logger *slog.Logger
	serv   *http.Server
}

func New(cfg app.Config, logger *slog.Logger, router *router.Router, options ...Option) *Server {
	serv := &http.Server{
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		Addr:              cfg.Addr(),
		Handler:           router.Setup(cfg),
	}

	for _, option := range options {
		option(serv)
	}

	return &Server{
		cfg:    cfg,
		logger: logger,
		serv:   serv,
	}
}

func (s *Server) Run() {
	s.logger.Info("Starting server",
		slog.String("app", s.cfg.Name),
		slog.String("addr", s.serv.Addr),
	)
	if err := s.serv.ListenAndServe(); err != nil {
		s.logger.Error(err.Error())
	}
}
