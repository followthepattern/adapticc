package auth

import (
	"context"
	"log/slog"

	"github.com/followthepattern/adapticc/config"
	"github.com/followthepattern/adapticc/container"
)

type AuthController struct {
	logger      *slog.Logger
	cfg         config.Config
	authService AuthenticationService
}

func NewAuthController(cont container.Container) AuthController {
	auth := NewAuthDatabase(cont.GetDB(), cont.GetLogger())
	authService := NewAuthenticationService(cont.GetConfig(), auth, cont.GetEmail(), cont.GetJWTKeys())

	return AuthController{
		authService: authService,
		cfg:         cont.GetConfig(),
		logger:      cont.GetLogger(),
	}
}

func (ctrl AuthController) Login(ctx context.Context, login LoginRequestParams) (*LoginResponse, error) {
	if err := login.Validate(); err != nil {
		return nil, err
	}

	return ctrl.authService.Login(ctx, login.Email, login.Password)
}

func (ctrl AuthController) Register(ctx context.Context, register RegisterRequestParams) (*RegisterResponse, error) {
	if err := register.Validate(); err != nil {
		return nil, err
	}

	return ctrl.authService.Register(ctx, register)
}
