package router

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"

	"github.com/seregaa020292/ModularMonolith/internal/config/app"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/errs"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/middleware"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/openapi"
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/server/respond"
	"github.com/seregaa020292/ModularMonolith/internal/ports/httprest"
)

type Router struct {
	swagger *openapi3.T
	openapi *httprest.OpenApiHandler
	appapi  *httprest.AppApiHandler
	respond *respond.Handle
	logger  *slog.Logger
}

func New(
	oapi *httprest.OpenApiHandler,
	appapi *httprest.AppApiHandler,
	respond *respond.Handle,
	logger *slog.Logger,
) (*Router, error) {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		return nil, err
	}

	return &Router{
		swagger: swagger,
		openapi: oapi,
		appapi:  appapi,
		respond: respond,
		logger:  logger,
	}, nil
}

func (router Router) Setup(cfg app.Config) http.Handler {
	r := chi.NewRouter()

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
	router.swagger.Servers = nil

	openapi.HandlerWithOptions(
		openapi.NewStrictHandlerWithOptions(router.openapi, nil, openapi.StrictHTTPServerOptions{
			RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				router.respond.Error(r.Context(), w, err)
			},
			ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				router.respond.Error(r.Context(), w, err)
			},
		}),
		openapi.ChiServerOptions{
			BaseRouter: r,
			Middlewares: []openapi.MiddlewareFunc{
				nethttpmiddleware.OapiRequestValidatorWithOptions(router.swagger, &nethttpmiddleware.Options{
					ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
						err := errs.NewBaseError(message, errors.New(http.StatusText(statusCode)), statusCode)
						router.respond.Error(middleware.LoggerWrapCtx(w), w, err)
					},
					Options: openapi3filter.Options{
						ExcludeRequestBody:        true,
						ExcludeRequestQueryParams: true,
					},
				}),
			},
			ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				router.respond.Error(r.Context(), w, err)
			},
		})

	r.Route("/admin", func(r chi.Router) {
		r.Use(chimiddleware.BasicAuth("Admin Panel", map[string]string{"admin": "admin"}))
		r.Get("/", router.appapi.AdminHandler.Home)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		router.respond.Error(r.Context(), w, errs.NewNotFoundError(nil))
	})

	return r
}
