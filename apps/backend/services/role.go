package services

import (
	"context"

	"github.com/followthepattern/adapticc/accesscontrol"
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/repositories/database"
	"github.com/followthepattern/adapticc/utils"
)

type Role struct {
	roleRepository database.Role
	ac             accesscontrol.AccessControl
}

func NewRole(cont container.Container, roleRepository database.Role) Role {
	return Role{
		roleRepository: roleRepository,
		ac:             cont.GetAccessControl().WithKind("role"),
	}
}

func (service Role) GetByID(ctx context.Context, id string) (*models.Role, error) {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return nil, err
	}

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.READ, id, roles...)
	if err != nil {
		return nil, err
	}

	result, err := service.roleRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service Role) Get(ctx context.Context, filter models.RoleListRequestParams) (*models.RoleListResponse, error) {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return nil, err
	}

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.READ, accesscontrol.ALLRESOURCE, roles...)
	if err != nil {
		return nil, err
	}

	return service.roleRepository.Get(filter)
}

func (service Role) AddRoleToUser(ctx context.Context, value models.UserRole) error {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return err
	}

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.CREATE, accesscontrol.NEW, roles...)
	if err != nil {
		return err
	}

	value.Userlog.CreationUserID = ctxu.ID

	return service.roleRepository.AddRoleToUser([]models.UserRole{value})
}

func (service Role) RemoveRoleFromUser(ctx context.Context, value models.UserRole) error {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return err
	}

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.CREATE, accesscontrol.NEW, roles...)
	if err != nil {
		return err
	}

	value.Userlog.CreationUserID = ctxu.ID

	return service.roleRepository.RemoveRoleFromUser(value)
}
