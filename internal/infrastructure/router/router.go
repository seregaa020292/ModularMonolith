package router

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"

	"github.com/seregaa020292/ModularMonolith/internal/config/consts"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/router/middleware"
	"github.com/seregaa020292/ModularMonolith/internal/ports/httprest"
)

type Router struct {
	mux     chi.Router
	rest    *httprest.HttpRest
	swagger *openapi3.T
}

func New(rest *httprest.HttpRest) (*Router, error) {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		return nil, err
	}

	return &Router{
		mux:     chi.NewRouter(),
		swagger: swagger,
		rest:    rest,
	}, nil
}

func (router Router) Setup() (http.Handler, error) {
	router.swagger.Servers = nil

	router.mux.Use(chimiddleware.Heartbeat("/health"))
	router.mux.Use(httprate.LimitByIP(consts.HttpRateRequestLimit, consts.HttpRateWindowLength))
	router.mux.Use(chimiddleware.StripSlashes)
	router.mux.Use(chimiddleware.RequestID)
	router.mux.Use(middleware.CorrelationID)
	router.mux.Use(chimiddleware.RealIP)
	router.mux.Use(chimiddleware.Recoverer)
	router.mux.Use(
		chimiddleware.SetHeader("X-Content-Type-Options", "nosniff"),
		chimiddleware.SetHeader("X-Frame-Options", "deny"),
		chimiddleware.SetHeader("X-Xss-Protection", "1; mode=block"),
	)

	router.mux.Group(func(r chi.Router) {
		r.Use(nethttpmiddleware.OapiRequestValidatorWithOptions(router.swagger, nil))
		openapi.HandlerFromMux(openapi.NewStrictHandler(router.rest, nil), r)
	})

	router.mux.Route("/admin", func(r chi.Router) {
		r.Use(chimiddleware.BasicAuth("Admin Panel", map[string]string{"admin": "admin"}))
		r.Get("/", router.rest.AdminHandler.Home)
	})

	router.mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
	})

	return router.mux, nil
}
