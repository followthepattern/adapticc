package controllers

import (
	"context"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/services"
)

type Auth struct {
	cfg         config.Config
	authService services.Auth
}

func NewAuth(cfg config.Config, authService services.Auth) Auth {
	return Auth{
		authService: authService,
		cfg:         cfg,
	}
}

func (ctrl Auth) Login(ctx context.Context, login models.LoginRequestParams) (*models.LoginResponse, error) {
	return ctrl.authService.Login(ctx, login.Email, login.Password)
}

func (ctrl Auth) Register(ctx context.Context, register models.RegisterRequestParams) (*models.RegisterResponse, error) {
	return ctrl.authService.Register(ctx, register)
}
