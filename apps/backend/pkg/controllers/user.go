package controllers

import (
	"context"
	"database/sql"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/services"
	"go.uber.org/zap"
)

type User struct {
	cfg         config.Config
	logger      *zap.Logger
	ctx         context.Context
	userService services.User
}

func NewUser(ctx context.Context, cerbosClient accesscontrol.AccessControl, db *sql.DB, cfg config.Config, logger *zap.Logger) User {
	userService := services.NewUser(ctx, cerbosClient, db, cfg, logger)

	return User{
		ctx:         ctx,
		cfg:         cfg,
		logger:      logger,
		userService: userService,
	}
}

func (ctrl User) GetByID(ctx context.Context, id string) (*models.User, error) {
	result, err := ctrl.userService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result.IsNil() {
		return nil, nil
	}

	return result, nil
}

func (ctrl User) Profile(ctx context.Context) (*models.User, error) {
	user, err := ctrl.userService.Profile(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ctrl User) Get(ctx context.Context, filter models.UserListRequestParams) (models.UserListResponse, error) {
	result, err := ctrl.userService.Get(ctx, filter)
	if err != nil {
		return models.UserListResponse{}, err
	}

	return result, nil
}

func (ctrl User) Create(ctx context.Context, user models.User) error {
	return ctrl.userService.Create(ctx, user)
}

func (ctrl User) Update(ctx context.Context, user models.User) error {
	return ctrl.userService.Update(ctx, user)
}

func (ctrl User) Delete(ctx context.Context, id string) error {
	return ctrl.userService.Delete(ctx, id)
}
