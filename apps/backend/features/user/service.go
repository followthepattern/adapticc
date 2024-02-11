package user

import (
	"context"

	"github.com/followthepattern/adapticc/accesscontrol"
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/types"
	"github.com/google/uuid"
)

type UserService struct {
	userRepository       UserDatabase
	authorizationService auth.AuthorizationService
}

func NewUserService(cont container.Container, authorizationService auth.AuthorizationService) UserService {
	repository := NewUserDatabase(cont.GetDB())

	user := UserService{
		userRepository:       repository,
		authorizationService: authorizationService,
	}

	return user
}

func (service UserService) GetByID(ctx context.Context, id types.String) (*UserModel, error) {
	err := service.authorizationService.Authorize(ctx, accesscontrol.READ, id.Data)
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
	ctxu, err := auth.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	return service.userRepository.GetByID(ctxu.ID)
}

func (service UserService) Get(ctx context.Context, request UserListRequestParams) (*UserListResponse, error) {
	err := service.authorizationService.Authorize(ctx, accesscontrol.READ, accesscontrol.ALLRESOURCE)
	if err != nil {
		return nil, err
	}

	request.Pagination.SetDefaultIfEmpty()

	return service.userRepository.Get(request)
}

func (service UserService) Create(ctx context.Context, value UserModel) error {
	userID, err := service.authorizationService.AuthorizedUser(ctx, accesscontrol.CREATE, accesscontrol.NEW)
	if err != nil {
		return err
	}

	value.ID = types.StringFrom(uuid.NewString())
	value.CreationUserID = types.StringFrom(userID)
	value.Active = types.FALSE

	return service.userRepository.Create([]UserModel{value})
}

func (service UserService) Update(ctx context.Context, value UserModel) error {
	userID, err := service.authorizationService.AuthorizedUser(ctx, accesscontrol.UPDATE, value.ID.Data)
	if err != nil {
		return err
	}

	value.UpdateUserID = types.StringFrom(userID)

	return service.userRepository.Update(value)
}

func (service UserService) ActivateUser(ctx context.Context, userID string) error {
	return service.userRepository.ActivateUser(userID)
}

func (service UserService) Delete(ctx context.Context, id types.String) error {
	err := service.authorizationService.Authorize(ctx, accesscontrol.DELETE, id.Data)
	if err != nil {
		return err
	}

	return service.userRepository.Delete(id)
}
