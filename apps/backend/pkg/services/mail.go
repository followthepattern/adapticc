package services

import (
	"net/smtp"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils"
)

type Mail struct {
	cfg   config.Mail
	email utils.Email
}

func NewMail(cfg config.Mail) Mail {
	return Mail{
		cfg:   cfg,
		email: utils.NewEmailWrapper(),
	}
}

func (service Mail) SendMail(mail models.Mail) error {
	config := service.cfg

	service.email.SetFrom(mail.From)
	service.email.SetTo(mail.To)
	service.email.SetSubject(mail.Subject)
	service.email.SetText(mail.Text)
	service.email.SetHTML(mail.HTML)

	err := service.email.
		Send(config.Addr,
			smtp.PlainAuth("",
				config.Username,
				config.Password,
				config.Host))

	return err
}
