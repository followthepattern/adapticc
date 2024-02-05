package controllers

import (
	"context"
	"log/slog"

	"github.com/followthepattern/adapticc/config"
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/repositories/database"
	"github.com/followthepattern/adapticc/services"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Role struct {
	roleService services.Role
	logger      *slog.Logger
	cfg         config.Config
}

func NewRole(cont container.Container) Role {
	roleRepository := database.NewRole(cont.GetDB())
	roleService := services.NewRole(cont, roleRepository)

	return Role{
		roleService: roleService,
		logger:      cont.GetLogger(),
		cfg:         cont.GetConfig(),
	}
}

func (ctrl Role) GetByID(ctx context.Context, id string) (*models.Role, error) {
	if err := validation.Validate(id, Required("productID")); err != nil {
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

func (ctrl Role) Get(ctx context.Context, filter models.RoleListRequestParams) (*models.RoleListResponse, error) {
	return ctrl.roleService.Get(ctx, filter)
}

func (ctrl Role) AddRoleToUser(ctx context.Context, value models.UserRole) error {
	if err := value.Validate(); err != nil {
		return err
	}

	return ctrl.roleService.AddRoleToUser(ctx, value)
}

func (ctrl Role) DeleteRoleFromUser(ctx context.Context, value models.UserRole) error {
	if err := value.Validate(); err != nil {
		return err
	}

	return ctrl.roleService.RemoveRoleFromUser(ctx, value)
}
