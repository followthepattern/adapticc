package controllers

import (
	"context"

	"log/slog"

	"github.com/followthepattern/adapticc/config"
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/services"
	"github.com/followthepattern/adapticc/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	cfg         config.Config
	logger      *slog.Logger
	userService services.User
}

func NewUser(cont container.Container) User {
	userService := services.NewUser(cont)

	return User{
		cfg:         cont.GetConfig(),
		logger:      cont.GetLogger(),
		userService: userService,
	}
}

func (ctrl User) GetByID(ctx context.Context, userID types.String) (*models.User, error) {
	if err := validation.Validate(userID, Required("userID")); err != nil {
		return nil, err
	}

	return ctrl.userService.GetByID(ctx, userID)
}

func (ctrl User) Profile(ctx context.Context) (*models.User, error) {
	user, err := ctrl.userService.Profile(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ctrl User) Get(ctx context.Context, filter models.UserListRequestParams) (*models.UserListResponse, error) {
	result, err := ctrl.userService.Get(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &result, nil
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

func (ctrl User) ActivateUser(ctx context.Context, userID string) error {
	if err := validation.Validate(userID, Required("userID")); err != nil {
		return err
	}

	return ctrl.userService.ActivateUser(ctx, userID)
}

func (ctrl User) Delete(ctx context.Context, userID types.String) error {
	if err := validation.Validate(userID, Required("userID")); err != nil {
		return err
	}

	return ctrl.userService.Delete(ctx, userID)
}
