package controllers

import (
	"context"
	"fmt"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
	"github.com/google/uuid"
)

type User struct {
	ctx        context.Context
	repository *database.User
}

func UserDependencyConstructor(cont *container.Container) (*User, error) {
	repository, err := container.Resolve[database.User](cont)
	if err != nil {
		return nil, err
	}

	return &User{
		ctx:        cont.GetContext(),
		repository: repository,
	}, nil
}

func (ctrl User) GetByID(ctx context.Context, id string) (*models.User, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return nil, fmt.Errorf("invalid user context")
	}

	result, err := ctrl.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if result.IsNil() {
		return nil, nil
	}

	return result, nil
}

func (ctrl User) Profile(ctx context.Context) (*models.User, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu.IsDefault() {
		return nil, fmt.Errorf("invalid user context")
	}

	user, err := ctrl.repository.GetByID(*ctxu.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ctrl User) Get(ctx context.Context, filter models.UserListRequestParams) (models.UserListResponse, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return models.UserListResponse{}, fmt.Errorf("invalid user context")
	}

	result, err := ctrl.repository.Get(*ctxu.ID, filter)
	if err != nil {
		return models.UserListResponse{}, err
	}

	return *result, nil
}

func (ctrl User) Create(ctx context.Context, user models.User) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	user.ID = pointers.ToPtr(uuid.New().String())

	return ctrl.repository.Create(*ctxu.ID, []models.User{user})
}

func (ctrl User) Update(ctx context.Context, user models.User) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	return ctrl.repository.Update(*ctxu.ID, user)
}

func (ctrl User) Delete(ctx context.Context, id string) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	return ctrl.repository.Delete(*ctxu.ID, id)
}
