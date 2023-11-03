package services

import (
	"context"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/utils"
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

func (service User) GetByID(ctx context.Context, id string) (*models.User, error) {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return nil, err
	}

	err = service.ac.Authorize(ctx, ctxu.ID, accesscontrol.READ, id, roles...)
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
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := service.userRepository.GetByID(ctxu.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service User) Get(ctx context.Context, filter models.UserListRequestParams) (models.UserListResponse, error) {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return models.UserListResponse{}, err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return models.UserListResponse{}, err
	}

	err = service.ac.Authorize(ctx, ctxu.ID, accesscontrol.READ, accesscontrol.ALLRESOURCE, roles...)
	if err != nil {
		return models.UserListResponse{}, err
	}

	result, err := service.userRepository.Get(filter)
	if err != nil {
		return models.UserListResponse{}, err
	}

	return *result, nil
}

func (service User) Create(ctx context.Context, value models.User) error {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return nil
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return err
	}

	err = service.ac.Authorize(ctx, ctxu.ID, accesscontrol.CREATE, accesscontrol.NEW, roles...)
	if err != nil {
		return err
	}

	value.ID = uuid.New().String()
	value.CreationUserID = ctxu.ID
	value.Active = false

	return service.userRepository.Create([]models.User{value})
}

func (service User) Update(ctx context.Context, value models.User) error {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return nil
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return err
	}

	err = service.ac.Authorize(ctx, ctxu.ID, accesscontrol.UPDATE, value.ID, roles...)
	if err != nil {
		return err
	}

	value.UpdateUserID = ctxu.ID

	return service.userRepository.Update(value)
}

func (service User) ActivateUser(ctx context.Context, userID string) error {
	return service.userRepository.ActivateUser(userID)
}

func (service User) Delete(ctx context.Context, id string) error {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return err
	}

	err = service.ac.Authorize(ctx, ctxu.ID, accesscontrol.DELETE, id, roles...)
	if err != nil {
		return err
	}

	return service.userRepository.Delete(id)
}
