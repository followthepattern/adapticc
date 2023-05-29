package rest

import (
	"encoding/json"
	"net/http"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/models"
	"go.uber.org/zap"
)

type mail struct {
	mail   *controllers.Mail
	logger *zap.Logger
}

func newMail(cont *container.Container) (*mail, error) {
	m, err := container.Resolve[controllers.Mail](cont)
	if err != nil {
		return nil, err
	}

	return &mail{
		mail:   m,
		logger: cont.GetLogger(),
	}, nil
}

func (m mail) SendGuestMessage(w http.ResponseWriter, r *http.Request) {
	var mail models.EmailSignIn
	err := json.NewDecoder(r.Body).Decode(&mail)
	if err != nil {
		BadRequest(w, "failed to decode request Body")
		return
	}

	if err = mail.Validate(); err != nil {
		BadRequest(w, err.Error())
		return
	}

	if err := m.mail.SendGuestMessage(r.Context(), mail); err != nil {
		m.logger.Error(err.Error())
		BadRequest(w, "failed to send the mail!")
		return
	}

	Success(w, "Email sent!")
}
