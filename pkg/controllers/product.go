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

type Product struct {
	repository *database.Product
}

func ProductDependencyConstructor(cont *container.Container) (*Product, error) {
	repository, err := container.Resolve[database.Product](cont)
	if err != nil {
		return nil, err
	}

	dependency := Product{
		repository: repository,
	}

	return &dependency, nil
}

func (service Product) GetByID(ctx context.Context, id string) (*models.Product, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return nil, fmt.Errorf("invalid user context")
	}

	result, err := service.repository.GetByID(*ctxu.ID, id)
	if err != nil {
		return nil, err
	}

	if result.IsNil() {
		return nil, nil
	}

	return result, nil
}

func (service Product) Get(ctx context.Context, filter models.ProductListRequestParams) (*models.ProductListResponse, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return nil, fmt.Errorf("invalid user context")
	}

	return service.repository.Get(*ctxu.ID, filter)
}

func (service Product) Create(ctx context.Context, value models.Product) error {
	if err := value.CreateValidate(); err != nil {
		return err
	}

	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	value.ID = pointers.ToPtr(uuid.New().String())

	return service.repository.Create(*ctxu.ID, []models.Product{value})
}

func (service Product) Update(ctx context.Context, value models.Product) error {
	if err := value.UpdateValidate(); err != nil {
		return err
	}

	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	return service.repository.Update(*ctxu.ID, value)
}

func (service Product) Delete(ctx context.Context, id string) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	return service.repository.Delete(*ctxu.ID, id)
}
