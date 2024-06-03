package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"

	"github.com/seregaa020292/ModularMonolith/internal/config/consts"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
	"github.com/seregaa020292/ModularMonolith/internal/ports/httprest"
)

type Router struct {
	mux  chi.Router
	rest *httprest.HttpRest
}

func NewRouter(rest *httprest.HttpRest) *Router {
	return &Router{
		mux:  chi.NewRouter(),
		rest: rest,
	}
}

func (router Router) Setup() (http.Handler, error) {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		return nil, err
	}

	swagger.Servers = nil

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

	router.mux.Group(func(r chi.Router) {
		r.Use(nethttpmiddleware.OapiRequestValidatorWithOptions(swagger, nil))
		openapi.HandlerFromMux(openapi.NewStrictHandler(router.rest, nil), r)
	})

	router.mux.Route("/admin", func(r chi.Router) {
		r.Use(middleware.BasicAuth("Admin Panel", map[string]string{"admin": "admin"}))
		r.Get("/", router.rest.AdminHandler.Home)
	})

	router.mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
	})

	return router.mux, nil
}
