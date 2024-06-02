package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"

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

func (router Router) Setup() http.Handler {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		panic(err)
	}

	swagger.Servers = nil

	router.setupMiddlewares()

	router.mux.Group(func(r chi.Router) {
		r.Use(nethttpmiddleware.OapiRequestValidatorWithOptions(swagger, nil))
		openapi.HandlerFromMux(openapi.NewStrictHandler(router.rest, nil), r)
	})

	router.mux.Route("/admin", func(r chi.Router) {
		r.Use(middleware.BasicAuth("Admin Panel", map[string]string{"admin": "admin"}))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello admin"))
		})
	})

	router.mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
	})

	return router.mux
}
