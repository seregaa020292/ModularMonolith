package router

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"

	"github.com/seregaa020292/ModularMonolith/internal/config"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/errs"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/http/middleware"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/http/response"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
	"github.com/seregaa020292/ModularMonolith/internal/ports/httprest"
)

type Router struct {
	mux     chi.Router
	swagger *openapi3.T
	rest    *httprest.ServerHandler
	errResp *response.ErrorHandle
	logger  *slog.Logger
}

func New(rest *httprest.ServerHandler, errResp *response.ErrorHandle, logger *slog.Logger) (*Router, error) {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		return nil, err
	}

	return &Router{
		mux:     chi.NewRouter(),
		swagger: swagger,
		rest:    rest,
		errResp: errResp,
		logger:  logger,
	}, nil
}

func (router Router) Setup(cfg config.App) http.Handler {
	router.swagger.Servers = nil
	r := router.mux

	r.Use(chimiddleware.Heartbeat("/health"))
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Use(chimiddleware.Timeout(60 * time.Second))
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.StripSlashes)
	r.Use(chimiddleware.RequestID)
	r.Use(middleware.CorrelationID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.RequestLogger(middleware.NewRequestLogger(router.logger)))
	r.Use(
		chimiddleware.SetHeader("X-Content-Type-Options", "nosniff"),
		chimiddleware.SetHeader("X-Frame-Options", "deny"),
		chimiddleware.SetHeader("X-Xss-Protection", "1; mode=block"),
	)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins(),
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Внедрение схемы OpenAPI.
	openapi.HandlerWithOptions(
		openapi.NewStrictHandlerWithOptions(router.rest.Openapi, nil, openapi.StrictHTTPServerOptions{
			RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				router.errResp.Send(r.Context(), w, err)
			},
			ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				router.errResp.Send(r.Context(), w, err)
			},
		}),
		openapi.ChiServerOptions{
			BaseRouter: r,
			Middlewares: []openapi.MiddlewareFunc{
				nethttpmiddleware.OapiRequestValidatorWithOptions(router.swagger, &nethttpmiddleware.Options{
					ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
						err := errs.NewBaseError(message, errors.New(http.StatusText(statusCode)), statusCode)
						router.errResp.Send(middleware.SetEntryLoggerCtxFromWriter(w), w, err)
					},
				}),
			},
			ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				router.errResp.Send(r.Context(), w, err)
			},
		})

	r.Route("/admin", func(r chi.Router) {
		r.Use(chimiddleware.BasicAuth("Admin Panel", map[string]string{"admin": "admin"}))
		r.Get("/", router.rest.App.AdminHandler.Home)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		router.errResp.Send(r.Context(), w, errs.NewNotFoundError(nil))
	})

	return r
}
