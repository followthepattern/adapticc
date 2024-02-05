package rest

import (
	"net/http"

	"github.com/followthepattern/adapticc/controllers"
	"github.com/go-chi/chi"
)

func New(ctrls controllers.Controllers) http.Handler {
	r := chi.NewMux()

	product := NewProduct(ctrls.Product())
	user := NewUser(ctrls.User())

	r.Route("/products", func(r chi.Router) {
		r.Post("/", product.Create)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/activate/{userID}", user.ActivateUser)
	})

	return r
}
