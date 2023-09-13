package rest

import (
	"net/http"

	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/go-chi/chi"
)

type RestConfig struct {
	product controllers.Product
}

func NewRestConfig(product controllers.Product) RestConfig {
	return RestConfig{
		product: product,
	}
}

func New(rest RestConfig) http.Handler {
	r := chi.NewMux()

	product := NewProduct(rest.product)

	r.Route("/products", func(r chi.Router) {
		r.Post("/", product.Create)
	})

	return r
}
