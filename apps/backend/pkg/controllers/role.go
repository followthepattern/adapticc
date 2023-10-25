package controllers

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/services"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Role struct {
	roleService services.Role
	logger      *slog.Logger
	cfg         config.Config
}

func NewRole(ctx context.Context, ac accesscontrol.AccessControl, db *sql.DB, cfg config.Config, logger *slog.Logger) Role {
	roleRepository := database.NewRole(ctx, db)
	roleService := services.NewRole(ctx, ac, roleRepository, logger)

	return Role{
		roleService: roleService,
		logger:      logger,
		cfg:         cfg,
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

	if result.IsNil() {
		return nil, nil
	}

	return result, nil
}

func (ctrl Role) Get(ctx context.Context, filter models.RoleListRequestParams) (*models.RoleListResponse, error) {
	return ctrl.roleService.Get(ctx, filter)
}
