package controllers

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/repositories/email"
	"github.com/followthepattern/adapticc/pkg/services"
)

type Auth struct {
	logger      *slog.Logger
	cfg         config.Config
	authService services.Auth
}

func NewAuth(ctx context.Context, db *sql.DB, emailClient email.Email, cfg config.Config, logger *slog.Logger) Auth {
	auth := database.NewAuth(ctx, db, logger)
	authService := services.NewAuth(cfg, auth, emailClient)

	return Auth{
		authService: authService,
		cfg:         cfg,
		logger:      logger,
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
