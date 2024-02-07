package services

import (
	"context"

	"github.com/followthepattern/adapticc/accesscontrol"
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/repositories/database"
	"github.com/followthepattern/adapticc/types"
	"github.com/google/uuid"
)

type User struct {
	userRepository database.User
	roleRepository database.Role
	ac             accesscontrol.AccessControl
}

func NewUser(cont container.Container) User {
	repository := database.NewUser(cont.GetDB())
	roleRepository := database.NewRole(cont.GetDB())

	user := User{
		userRepository: repository,
		roleRepository: roleRepository,
		ac:             cont.GetAccessControl().WithKind("user"),
	}

	return user
}

func (service User) GetByID(ctx context.Context, id types.String) (*models.User, error) {
	ctxu, err := GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return nil, err
	}

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.READ, id.Data, roles...)
	if err != nil {
		return nil, err
	}

	result, err := service.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if result.IsDefault() {
		return nil, nil
	}

	return result, nil
}

func (service User) Profile(ctx context.Context) (*models.User, error) {
	ctxu, err := GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	return service.userRepository.GetByID(ctxu.ID)
}

func (service User) Get(ctx context.Context, request models.UserListRequestParams) (models.UserListResponse, error) {
	ctxu, err := GetUserContext(ctx)
	if err != nil {
		return models.UserListResponse{}, err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return models.UserListResponse{}, err
	}

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.READ, accesscontrol.ALLRESOURCE, roles...)
	if err != nil {
		return models.UserListResponse{}, err
	}

	request.Pagination.SetDefaultIfEmpty()

	result, err := service.userRepository.Get(request)
	if err != nil {
		return models.UserListResponse{}, err
	}

	return *result, nil
}

func (service User) Create(ctx context.Context, value models.User) error {
	ctxu, err := GetUserContext(ctx)
	if err != nil {
		return nil
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return err
	}

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.CREATE, accesscontrol.NEW, roles...)
	if err != nil {
		return err
	}

	value.ID = types.StringFrom(uuid.NewString())
	value.CreationUserID = ctxu.ID
	value.Active = types.FALSE

	return service.userRepository.Create([]models.User{value})
}

func (service User) Update(ctx context.Context, value models.User) error {
	ctxu, err := GetUserContext(ctx)
	if err != nil {
		return nil
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return err
	}

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.UPDATE, value.ID.Data, roles...)
	if err != nil {
		return err
	}

	value.UpdateUserID = ctxu.ID

	return service.userRepository.Update(value)
}

func (service User) ActivateUser(ctx context.Context, userID string) error {
	return service.userRepository.ActivateUser(userID)
}

func (service User) Delete(ctx context.Context, id types.String) error {
	ctxu, err := GetUserContext(ctx)
	if err != nil {
		return err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return err
	}

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.DELETE, id.Data, roles...)
	if err != nil {
		return err
	}

	return service.userRepository.Delete(id)
}
