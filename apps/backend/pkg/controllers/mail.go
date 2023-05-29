package controllers

import (
	"bytes"
	"context"
	"html/template"
	"time"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/request"
	"github.com/followthepattern/adapticc/pkg/services"
)

const (
	FPInfoMailAddress = "Follow The Pattern <info@followthepattern.net>"
	AdminEmail        = "csaba@followthepattern.net"
	TemplatePath      = "./templates/guest_mail.tmpl"
)

type Mail struct {
	mailMsgChannelOut chan<- request.RequestHandler[models.Mail, struct{}]
	sendMsg           func(context.Context, request.RequestHandler[models.Mail, struct{}]) error
}

func MailDependencyConstructor(cont *container.Container) (*Mail, error) {
	mailMsgChannelOut, err := container.Resolve[services.MailMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	dependency := Mail{
		mailMsgChannelOut: *mailMsgChannelOut,
		sendMsg:           request.CreateSenderFunc(*mailMsgChannelOut, request.DefaultTimeOut),
	}

	return &dependency, nil
}

func (ctrl Mail) SendGuestMessage(ctx context.Context, message models.EmailSignIn) error {
	tmpl, err := template.New("guest_mail.tmpl").ParseFiles(TemplatePath)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, &message)
	if err != nil {
		return err
	}

	mail := models.Mail{
		From:    FPInfoMailAddress,
		To:      []string{AdminEmail},
		Subject: "FP New message",
		HTML:    b.Bytes(),
	}

	customTimeout := request.TimeoutOption[models.Mail, struct{}](time.Second * 5)

	req := request.New(
		ctx,
		mail,
		customTimeout)

	if err := ctrl.sendMsg(ctx, req); err != nil {
		return err
	}

	_, err = req.Wait()
	return err
}
