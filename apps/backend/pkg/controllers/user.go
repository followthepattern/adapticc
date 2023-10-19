package controllers

import (
	"context"
	"database/sql"

	"log/slog"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/services"
	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	cfg         config.Config
	logger      *slog.Logger
	ctx         context.Context
	userService services.User
}

func NewUser(ctx context.Context, cerbosClient accesscontrol.AccessControl, db *sql.DB, cfg config.Config, logger *slog.Logger) User {
	userService := services.NewUser(ctx, cerbosClient, db, cfg, logger)

	return User{
		ctx:         ctx,
		cfg:         cfg,
		logger:      logger,
		userService: userService,
	}
}

func (ctrl User) GetByID(ctx context.Context, id string) (*models.User, error) {
	if err := validation.Validate(id, Required("userID")); err != nil {
		return nil, err
	}

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

func (ctrl User) Create(ctx context.Context, value models.User) error {
	if err := value.CreateValidate(); err != nil {
		return err
	}

	return ctrl.userService.Create(ctx, value)
}

func (ctrl User) Update(ctx context.Context, value models.User) error {
	if err := value.UpdateValidate(); err != nil {
		return err
	}

	return ctrl.userService.Update(ctx, value)
}

func (ctrl User) Delete(ctx context.Context, id string) error {
	if err := validation.Validate(id, Required("userID")); err != nil {
		return err
	}

	return ctrl.userService.Delete(ctx, id)
}
