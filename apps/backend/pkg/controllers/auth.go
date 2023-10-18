package controllers

import (
	"context"
	"database/sql"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/services"
	"go.uber.org/zap"
)

type Auth struct {
	logger      *zap.Logger
	cfg         config.Config
	authService services.Auth
}

func NewAuth(ctx context.Context, db *sql.DB, cfg config.Config, logger *zap.Logger) Auth {
	auth := database.NewAuth(ctx, db, logger)
	authService := services.NewAuth(cfg, auth)

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
