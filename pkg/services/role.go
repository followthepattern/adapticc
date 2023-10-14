package services

import (
	"context"
	"database/sql"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/utils"
	"go.uber.org/zap"
)

type Role struct {
	roleRepository database.Role
	ac             accesscontrol.AccessControl
}

func NewRole(ctx context.Context, ac accesscontrol.AccessControl, db *sql.DB, cfg config.Config, logger *zap.Logger) User {
	repository := database.NewUser(ctx, db)
	roleRepository := database.NewRole(ctx, db)

	user := User{
		userRepository: repository,
		roleRepository: roleRepository,
		ac:             ac,
	}

	return user
}

func (service Role) GetByID(ctx context.Context, id string) (*models.Role, error) {
	ctxu, err := utils.GetUserContext(ctx)
	if err == nil {
		return nil, err
	}

	roles, err := service.roleRepository.GetProfileRolesArray(*ctxu.ID)
	if err != nil {
		return nil, err
	}

	err = service.ac.Authorize(ctx, *ctxu.ID, accesscontrol.READ, id, roles...)
	if err != nil {
		return nil, err
	}

	result, err := service.roleRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
