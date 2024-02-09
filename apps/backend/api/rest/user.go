package rest

import (
	"net/http"

	"github.com/followthepattern/adapticc/user"
	"github.com/go-chi/chi"
)

type User struct {
	user user.UserController
}

func NewUser(ctrl user.UserController) User {
	return User{
		user: ctrl,
	}
}

func (rest User) ActivateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	if err := rest.user.ActivateUser(r.Context(), userID); err != nil {
		BadRequest(w, err.Error())
		return
	}

	Success(w)
}
