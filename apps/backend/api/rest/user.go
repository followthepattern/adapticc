package rest

import (
	"net/http"

	"github.com/followthepattern/adapticc/controllers"
	"github.com/go-chi/chi"
)

type User struct {
	user controllers.User
}

func NewUser(ctrl controllers.User) User {
	return User{
		user: ctrl,
	}
}

func (service User) ActivateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	if err := service.user.ActivateUser(r.Context(), userID); err != nil {
		BadRequest(w, err.Error())
		return
	}

	Success(w)
}
