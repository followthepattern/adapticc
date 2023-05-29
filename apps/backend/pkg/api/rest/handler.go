package rest

import (
	"net/http"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/go-chi/chi"
)

func NewHandler(cont *container.Container) (http.Handler, error) {
	r := chi.NewMux()

	mail, err := newMail(cont)
	if err != nil {
		return nil, err
	}

	r.Route("/mail", func(r chi.Router) {
		r.Post("/sendguestmessage", mail.SendGuestMessage)
	})
	return r, nil
}
