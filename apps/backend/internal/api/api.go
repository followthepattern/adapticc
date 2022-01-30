package api

import (
	"backend/internal/api/graphql_api"
	"backend/internal/api/middlewares"
	"backend/internal/container"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func GetRouter(cont container.IContainer) (*chi.Mux, error) {
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
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/healthcheck"))

	handler, err := graphql_api.NewHandler(cont)

	authMiddleware := middlewares.NewAuth(cont)

	if err != nil {
		return nil, err
	}

	r.Route("/", func(r chi.Router) {
		r.With(authMiddleware.Authenticate).Post("/graphql", handler.ServeHTTP)
	})

	return r, nil
}
