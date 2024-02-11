package user

import (
	"net/http"

	"github.com/followthepattern/adapticc/api"
	"github.com/go-chi/chi"
)

type UserRest struct {
	user UserController
}

func NewUserRest(ctrl UserController) UserRest {
	return UserRest{
		user: ctrl,
	}
}

func (rest UserRest) ActivateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	if err := rest.user.ActivateUser(r.Context(), userID); err != nil {
		api.BadRequest(w, err.Error())
		return
	}

	api.Success(w)
}
