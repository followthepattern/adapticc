package controllers

import (
	"context"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/request"
	"github.com/followthepattern/adapticc/pkg/services"
)

type Auth struct {
	cfg               config.Config
	authMsgChannelOut chan<- models.AuthMsg
	sendMsg           func(ctx context.Context, msg models.AuthMsg) error
}

func AuthDependencyConstructor(cont *container.Container) (*Auth, error) {
	authMsgChannelOut, err := container.Resolve[services.AuthMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	dependency := Auth{
		authMsgChannelOut: *authMsgChannelOut,
		sendMsg:           request.CreateSenderFunc(*authMsgChannelOut, request.DefaultTimeOut),
	}

	dependency.cfg = cont.GetConfig()

	return &dependency, nil
}

func (ctrl Auth) Login(ctx context.Context, login models.LoginRequestParams) (*models.LoginResponse, error) {
	req := request.New[models.LoginRequestParams, models.LoginResponse](ctx, login)

	msg := models.AuthMsg{
		Login: &req,
	}

	if err := ctrl.sendMsg(ctx, msg); err != nil {
		return nil, err
	}

	resp, err := req.Wait()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ctrl Auth) Register(ctx context.Context, register models.RegisterRequestParams) (*models.RegisterResponse, error) {
	req := request.New[models.RegisterRequestParams, models.RegisterResponse](ctx, register)

	msg := models.AuthMsg{
		Register: &req,
	}

	if err := ctrl.sendMsg(ctx, msg); err != nil {
		return nil, err
	}

	resp, err := req.Wait()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
