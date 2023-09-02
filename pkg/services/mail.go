package services

import (
	"context"
	"net/smtp"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils"
)

type Mail struct {
	cfg   config.Mail
	ctx   context.Context
	email utils.Email
}

func MailDependencyConstructor(cont *container.Container) (*Mail, error) {
	config := cont.GetConfig().Mail

	if err := config.Validate(); err != nil {
		return nil, err
	}

	dependency := Mail{
		ctx:   cont.GetContext(),
		email: utils.NewEmailWrapper(),
	}

	return &dependency, nil
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
