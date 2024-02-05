package api

import (
	"net/http"

	"log/slog"

	"github.com/followthepattern/adapticc/api/middlewares"
	"github.com/followthepattern/adapticc/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func NewHttpApi(cfg config.Config,
	jwtKeys config.JwtKeyPair,
	graphqlHandler http.Handler,
	restHandler http.Handler,
	logger *slog.Logger,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middlewares.SessionContextID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middlewares.Heartbeat("/healthcheck", cfg.Version))

	middlewares.AddMiddlewareLogger(r, logger)

	authMiddleware := middlewares.NewJWT(logger, jwtKeys)

	r.Route("/", func(r chi.Router) {
		r.With(authMiddleware.Authenticate).Post("/graphql", graphqlHandler.ServeHTTP)
		r.With(authMiddleware.Authenticate).Mount("/", restHandler)
	})

	return r
}
