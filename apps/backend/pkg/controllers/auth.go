package controllers

import (
	"context"
	"log/slog"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/services"
)

type Auth struct {
	logger      *slog.Logger
	cfg         config.Config
	authService services.Auth
}

func NewAuth(cont container.Container) Auth {
	auth := database.NewAuth(cont.GetDB(), cont.GetLogger())
	authService := services.NewAuth(cont.GetConfig(), auth, cont.GetEmail(), cont.GetJWTKeys())

	return Auth{
		authService: authService,
		cfg:         cont.GetConfig(),
		logger:      cont.GetLogger(),
	}
}

func (ctrl Auth) Login(ctx context.Context, login models.LoginRequestParams) (*models.LoginResponse, error) {
	if err := login.Validate(); err != nil {
		return nil, err
	}

	return ctrl.authService.Login(ctx, login.Email, login.Password)
}

func (ctrl Auth) Register(ctx context.Context, register models.RegisterRequestParams) (*models.RegisterResponse, error) {
	if err := register.Validate(); err != nil {
		return nil, err
	}

	return ctrl.authService.Register(ctx, register)
}
