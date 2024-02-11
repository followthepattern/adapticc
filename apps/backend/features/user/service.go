package user

import (
	"context"

	"github.com/followthepattern/adapticc/accesscontrol"
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/services"
	"github.com/followthepattern/adapticc/types"
	"github.com/google/uuid"
)

type UserService struct {
	userRepository UserDatabase
	ac             accesscontrol.AccessControl
}

func NewUserService(cont container.Container) UserService {
	repository := NewUserDatabase(cont.GetDB())

	user := UserService{
		userRepository: repository,
		ac:             cont.GetAccessControl().WithKind("user"),
	}

	return user
}

func (service UserService) GetByID(ctx context.Context, id types.String) (*UserModel, error) {
	ctxu, err := services.GetUserContext(ctx)
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

func (service UserService) Profile(ctx context.Context) (*UserModel, error) {
	ctxu, err := GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	return service.userRepository.GetByID(ctxu.ID)
}

func (service UserService) Get(ctx context.Context, request UserListRequestParams) (UserListResponse, error) {
	ctxu, err := GetUserContext(ctx)
	if err != nil {
		return UserListResponse{}, err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return UserListResponse{}, err
	}

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.READ, accesscontrol.ALLRESOURCE, roles...)
	if err != nil {
		return UserListResponse{}, err
	}

	request.Pagination.SetDefaultIfEmpty()

	result, err := service.userRepository.Get(request)
	if err != nil {
		return UserListResponse{}, err
	}

	return *result, nil
}

func (service UserService) Create(ctx context.Context, value UserModel) error {
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

	return service.userRepository.Create([]UserModel{value})
}

func (service UserService) Update(ctx context.Context, value UserModel) error {
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

func (service UserService) ActivateUser(ctx context.Context, userID string) error {
	return service.userRepository.ActivateUser(userID)
}

func (service UserService) Delete(ctx context.Context, id types.String) error {
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
