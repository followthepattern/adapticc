package role

import (
	"context"
	"log/slog"

	"github.com/followthepattern/adapticc/config"
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/types"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RoleController struct {
	roleService RoleService
	logger      *slog.Logger
	cfg         config.Config
}

func NewRoleController(cont container.Container) RoleController {
	roleRepository := NewRoleDatabase(cont.GetDB())
	authorizationService := auth.NewAuthorizationService(cont, "role")
	roleService := NewRoleService(cont, authorizationService, roleRepository)

	return RoleController{
		roleService: roleService,
		logger:      cont.GetLogger(),
		cfg:         cont.GetConfig(),
	}
}

func (ctrl RoleController) GetByID(ctx context.Context, id string) (*RoleModel, error) {
	if err := validation.Validate(id, types.Required("productID")); err != nil {
		return nil, err
	}

	result, err := ctrl.roleService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result.IsDefault() {
		return nil, nil
	}

	return result, nil
}

func (ctrl RoleController) Get(ctx context.Context, filter RoleListRequestParams) (*RoleListResponse, error) {
	return ctrl.roleService.Get(ctx, filter)
}

func (ctrl RoleController) AddRoleToUser(ctx context.Context, value UserRoleModel) error {
	if err := value.Validate(); err != nil {
		return err
	}

	return ctrl.roleService.AddRoleToUser(ctx, value)
}

func (ctrl RoleController) DeleteRoleFromUser(ctx context.Context, value UserRoleModel) error {
	if err := value.Validate(); err != nil {
		return err
	}

	return ctrl.roleService.RemoveRoleFromUser(ctx, value)
}
