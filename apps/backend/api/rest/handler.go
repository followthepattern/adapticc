package rest

import (
	"net/http"

	"github.com/followthepattern/adapticc/controllers"
	"github.com/followthepattern/adapticc/features/product"
	"github.com/followthepattern/adapticc/features/user"
	"github.com/go-chi/chi"
)

func New(ctrls controllers.Controllers) http.Handler {
	r := chi.NewMux()

	product := product.NewProductRest(ctrls.Product())
	user := user.NewUserRest(ctrls.User())

	r.Route("/products", func(r chi.Router) {
		r.Post("/", product.Create)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/activate/{userID}", user.ActivateUser)
	})

	return r
}
