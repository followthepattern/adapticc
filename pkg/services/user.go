package services

import (
	"context"
	"database/sql"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type User struct {
	userRepository database.User
	roleRepository database.Role
	ac             accesscontrol.AccessControl
}

func NewUser(ctx context.Context, ac accesscontrol.AccessControl, db *sql.DB, cfg config.Config, logger *zap.Logger) User {
	repository := database.NewUser(ctx, db)
	roleRepository := database.NewRole(ctx, db)

	user := User{
		userRepository: repository,
		roleRepository: roleRepository,
		ac:             ac.WithKind("user"),
	}

	return user
}

func (service User) GetByID(ctx context.Context, id string) (*models.User, error) {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
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

	result, err := service.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if result.IsNil() {
		return nil, nil
	}

	return result, nil
}

func (service User) Profile(ctx context.Context) (*models.User, error) {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := service.userRepository.GetByID(*ctxu.ID)
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

	roles, err := service.roleRepository.GetProfileRolesArray(*ctxu.ID)
	if err != nil {
		return models.UserListResponse{}, err
	}

	err = service.ac.Authorize(ctx, *ctxu.ID, accesscontrol.READ, accesscontrol.ALLRESOURCE, roles...)
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

	roles, err := service.roleRepository.GetProfileRolesArray(*ctxu.ID)
	if err != nil {
		return err
	}

	err = service.ac.Authorize(ctx, *ctxu.ID, accesscontrol.CREATE, accesscontrol.NEW, roles...)
	if err != nil {
		return err
	}

	value.ID = pointers.ToPtr(uuid.New().String())
	value.CreationUserID = ctxu.ID
	value.Active = pointers.ToPtr(false)

	return service.userRepository.Create([]models.User{value})
}

func (service User) Update(ctx context.Context, value models.User) error {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return nil
	}

	roles, err := service.roleRepository.GetProfileRolesArray(*ctxu.ID)
	if err != nil {
		return err
	}

	err = service.ac.Authorize(ctx, *ctxu.ID, accesscontrol.UPDATE, *value.ID, roles...)
	if err != nil {
		return err
	}

	return service.userRepository.Update(*ctxu.ID, value)
}

func (service User) Delete(ctx context.Context, id string) error {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return err
	}

	roles, err := service.roleRepository.GetProfileRolesArray(*ctxu.ID)
	if err != nil {
		return err
	}

	err = service.ac.Authorize(ctx, *ctxu.ID, accesscontrol.DELETE, id, roles...)
	if err != nil {
		return err
	}

	return service.userRepository.Delete(*ctxu.ID, id)
}
