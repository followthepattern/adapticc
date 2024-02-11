package role

import (
	"context"

	"github.com/followthepattern/adapticc/accesscontrol"
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/types"
)

type RoleService struct {
	roleRepository RoleDatabase
	authorization  auth.AuthorizationService
}

func NewRoleService(
	cont container.Container,
	authorization auth.AuthorizationService) RoleService {
	roleRepository := NewRoleDatabase(cont.GetDB())
	return RoleService{
		roleRepository: roleRepository,
		authorization:  authorization,
	}
}

func (service RoleService) GetByID(ctx context.Context, id string) (*RoleModel, error) {
	err := service.authorization.Authorize(ctx, accesscontrol.READ, id)
	if err != nil {
		return nil, err
	}

	result, err := service.roleRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service RoleService) Get(ctx context.Context, request RoleListRequestParams) (*RoleListResponse, error) {
	err := service.authorization.Authorize(ctx, accesscontrol.READ, accesscontrol.ALLRESOURCE)
	if err != nil {
		return nil, err
	}

	request.Pagination.SetDefaultIfEmpty()

	return service.roleRepository.Get(request)
}

func (service RoleService) AddRoleToUser(ctx context.Context, value UserRoleModel) error {
	userID, err := service.authorization.AuthorizedUser(ctx, accesscontrol.CREATE, accesscontrol.NEW)
	if err != nil {
		return err
	}

	value.Userlog.CreationUserID = types.StringFrom(userID)

	return service.roleRepository.AddRoleToUser([]UserRoleModel{value})
}

func (service RoleService) RemoveRoleFromUser(ctx context.Context, value UserRoleModel) error {
	err := service.authorization.Authorize(ctx, accesscontrol.DELETE, accesscontrol.ALLRESOURCE)
	if err != nil {
		return err
	}

	return service.roleRepository.RemoveRoleFromUser(value)
}
