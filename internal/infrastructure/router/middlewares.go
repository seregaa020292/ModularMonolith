package router

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"

	"github.com/seregaa020292/ModularMonolith/internal/config/consts"
)

func (router Router) setupMiddlewares() {
	router.mux.Use(middleware.Heartbeat("/health"))
	router.mux.Use(httprate.LimitByIP(consts.HttpRateRequestLimit, consts.HttpRateWindowLength))
	router.mux.Use(middleware.StripSlashes)
	router.mux.Use(middleware.RequestID)
	router.mux.Use(middleware.RealIP)
	router.mux.Use(middleware.Recoverer)
	router.mux.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
		middleware.SetHeader("X-Xss-Protection", "1; mode=block"),
	)
}
