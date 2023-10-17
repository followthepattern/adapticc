package rest

import (
	"net/http"

	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/go-chi/chi"
)

func New(ctrls controllers.Controllers) http.Handler {
	r := chi.NewMux()

	product := NewProduct(ctrls.Product())

	r.Route("/products", func(r chi.Router) {
		r.Post("/", product.Create)
	})

	return r
}
