package user

import (
	"context"

	"log/slog"

	"github.com/followthepattern/adapticc/config"
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserController struct {
	cfg         config.Config
	logger      *slog.Logger
	userService UserService
}

func NewUserController(cont container.Container) UserController {
	authorizationService := auth.NewAuthorizationService(cont, "user")
	userService := NewUserService(cont, authorizationService)

	return UserController{
		cfg:         cont.GetConfig(),
		logger:      cont.GetLogger(),
		userService: userService,
	}
}

func (ctrl UserController) GetByID(ctx context.Context, userID types.String) (*UserModel, error) {
	if err := validation.Validate(userID, types.Required("userID")); err != nil {
		return nil, err
	}

	return ctrl.userService.GetByID(ctx, userID)
}

func (ctrl UserController) Profile(ctx context.Context) (*UserModel, error) {
	user, err := ctrl.userService.Profile(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ctrl UserController) Get(ctx context.Context, filter UserListRequestParams) (*UserListResponse, error) {
	return ctrl.userService.Get(ctx, filter)
}

func (ctrl UserController) Create(ctx context.Context, value UserModel) error {
	if err := value.CreateValidate(); err != nil {
		return err
	}

	return ctrl.userService.Create(ctx, value)
}

func (ctrl UserController) Update(ctx context.Context, value UserModel) error {
	if err := value.UpdateValidate(); err != nil {
		return err
	}

	return ctrl.userService.Update(ctx, value)
}

func (ctrl UserController) ActivateUser(ctx context.Context, userID string) error {
	if err := validation.Validate(userID, types.Required("userID")); err != nil {
		return err
	}

	return ctrl.userService.ActivateUser(ctx, userID)
}

func (ctrl UserController) Delete(ctx context.Context, userID types.String) error {
	if err := validation.Validate(userID, types.Required("userID")); err != nil {
		return err
	}

	return ctrl.userService.Delete(ctx, userID)
}
