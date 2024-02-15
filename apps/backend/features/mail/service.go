package mail

import (
	"net/smtp"

	"github.com/followthepattern/adapticc/config"
)

type Mail struct {
	cfg   config.Mail
	email Email
}

func NewMail(cfg config.Mail, emailClient Email) Mail {
	return Mail{
		cfg:   cfg,
		email: emailClient,
	}
}

func (service Mail) SendMail(mail MailModel) error {
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
