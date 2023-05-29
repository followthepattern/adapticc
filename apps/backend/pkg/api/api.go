package api

import (
	"github.com/followthepattern/adapticc/pkg/api/graphql"
	"github.com/followthepattern/adapticc/pkg/api/middlewares"
	"github.com/followthepattern/adapticc/pkg/api/rest"
	"github.com/followthepattern/adapticc/pkg/container"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func GetRouter(cont *container.Container) (*chi.Mux, error) {
	r := chi.NewRouter()

	cfg := cont.GetConfig()

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
	r.Use(middlewares.Heartbeat("/healthcheck", cfg.Version))

	graphqlHandler, err := graphql.NewHandler(cont)
	if err != nil {
		return nil, err
	}
	restHandler, err := rest.NewHandler(cont)
	if err != nil {
		return nil, err
	}

	authMiddleware := middlewares.NewJWT(cont)

	r.Route("/", func(r chi.Router) {
		r.With(authMiddleware.Authenticate).Post("/graphql", graphqlHandler.ServeHTTP)
		r.Mount("/", restHandler)
	})

	return r, nil
}
