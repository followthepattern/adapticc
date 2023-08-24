package rest

import (
	"net/http"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/go-chi/chi"
)

func NewHandler(cont *container.Container) (http.Handler, error) {
	r := chi.NewMux()

	product, err := newProduct(cont)
	if err != nil {
		return nil, err
	}

	r.Route("/products", func(r chi.Router) {
		r.Post("/", product.Create)
	})

	return r, nil
}
