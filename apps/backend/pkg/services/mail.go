package services

import (
	"context"
	"net/smtp"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/request"
	"github.com/followthepattern/adapticc/pkg/utils"
)

type MailMsgChannel chan request.RequestHandler[models.Mail, struct{}]

func RegisterMailChannel(cont *container.Container) {
	if cont == nil {
		return
	}
	mailMsgChannel := make(MailMsgChannel)
	container.Register(cont, func(cont *container.Container) (*MailMsgChannel, error) {
		return &mailMsgChannel, nil
	})
}

type Mail struct {
	cfg   config.Mail
	ctx   context.Context
	email utils.Email

	mailMsgChannel <-chan request.RequestHandler[models.Mail, struct{}]
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

	mailMsgChannel, err := container.Resolve[MailMsgChannel](cont)
	if err != nil {
		return nil, err
	}
	dependency.mailMsgChannel = *mailMsgChannel

	go func() {
		dependency.MonitorChannels()
	}()

	return &dependency, nil
}

func (service Mail) MonitorChannels() {
	for {
		select {
		case req := <-service.mailMsgChannel:
			service.replyRequest(req)
		case <-service.ctx.Done():
			return
		}
	}
}

func (service Mail) replyRequest(req request.RequestHandler[models.Mail, struct{}]) {
	requestParams := req.RequestParams()
	if err := service.sendMail(requestParams); err != nil {
		req.ReplyError(err)
		return
	}
	req.Reply(request.Success())
}

func (service Mail) sendMail(mail models.Mail) error {
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
