package router

import (
	"errors"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"

	"github.com/seregaa020292/ModularMonolith/internal/config/consts"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/errs"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/http/middleware"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/http/response"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
	"github.com/seregaa020292/ModularMonolith/internal/ports/httprest"
)

type Router struct {
	mux     chi.Router
	swagger *openapi3.T
	rest    *httprest.HttpRest
	errResp *response.ErrorResponse
}

func New(rest *httprest.HttpRest, errResp *response.ErrorResponse) (*Router, error) {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		return nil, err
	}

	return &Router{
		mux:     chi.NewRouter(),
		swagger: swagger,
		rest:    rest,
		errResp: errResp,
	}, nil
}

func (router Router) Setup() (http.Handler, error) {
	router.swagger.Servers = nil
	r := router.mux

	r.Use(chimiddleware.Heartbeat("/health"))
	r.Use(httprate.LimitByIP(consts.HttpRateRequestLimit, consts.HttpRateWindowLength))
	r.Use(chimiddleware.StripSlashes)
	r.Use(chimiddleware.RequestID)
	r.Use(middleware.CorrelationID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Recoverer)
	r.Use(
		chimiddleware.SetHeader("X-Content-Type-Options", "nosniff"),
		chimiddleware.SetHeader("X-Frame-Options", "deny"),
		chimiddleware.SetHeader("X-Xss-Protection", "1; mode=block"),
	)

	r.Group(func(r chi.Router) {
		r.Use(nethttpmiddleware.OapiRequestValidatorWithOptions(router.swagger, &nethttpmiddleware.Options{
			ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
				router.errResp.Send(w,
					errs.NewBaseError(message, errors.New(http.StatusText(statusCode)), statusCode))
			},
		}))
		openapi.HandlerFromMux(openapi.NewStrictHandlerWithOptions(router.rest, nil,
			openapi.StrictHTTPServerOptions{
				RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
					router.errResp.Send(w, err)
				},
				ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
					router.errResp.Send(w, err)
				},
			}), r)
	})

	r.Route("/admin", func(r chi.Router) {
		r.Use(chimiddleware.BasicAuth("Admin Panel", map[string]string{"admin": "admin"}))
		r.Get("/", router.rest.AdminHandler.Home)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
	})

	return r, nil
}
