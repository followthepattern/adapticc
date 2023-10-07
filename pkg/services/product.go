package services

import (
	"context"
	"errors"
	"fmt"

	cerbos "github.com/cerbos/cerbos/client"
	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Product struct {
	repository   database.Product
	logger       *zap.Logger
	cfg          config.Config
	cerbosClient cerbos.Client
}

func (Product) resourceName() string {
	return "product"
}

func NewProduct(ctx context.Context, cerbosClient cerbos.Client, productRepository database.Product, cfg config.Config, logger *zap.Logger) Product {
	return Product{
		repository:   productRepository,
		logger:       logger,
		cfg:          cfg,
		cerbosClient: cerbosClient,
	}
}

func (service Product) IsAllowed(ctx context.Context, action string, resourceID string) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)

	principal := cerbos.NewPrincipal(*ctxu.ID, "product:editor")

	resource := cerbos.NewResource(service.resourceName(), resourceID)

	ok, err := service.cerbosClient.IsAllowed(ctx, principal, resource, action)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("action is not allowed")
	}

	return nil
}

func (service Product) GetByID(ctx context.Context, id string) (*models.Product, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return nil, fmt.Errorf("invalid user context")
	}

	err := service.IsAllowed(ctx, accesscontrol.READ, id)
	if err != nil {
		return nil, err
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

	err := service.IsAllowed(ctx, accesscontrol.READ, "all")
	if err != nil {
		return nil, err
	}

	return service.repository.Get(*ctxu.ID, filter)
}

func (service Product) Create(ctx context.Context, value models.Product) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	value.ID = pointers.ToPtr(uuid.New().String())

	err := service.IsAllowed(ctx, accesscontrol.CREATE, *value.ID)
	if err != nil {
		return err
	}

	return service.repository.Create(*ctxu.ID, []models.Product{value})
}

func (service Product) Update(ctx context.Context, value models.Product) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	err := service.IsAllowed(ctx, accesscontrol.UPDATE, *value.ID)
	if err != nil {
		return err
	}

	return service.repository.Update(*ctxu.ID, value)
}

func (service Product) Delete(ctx context.Context, id string) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	err := service.IsAllowed(ctx, accesscontrol.DELETE, id)
	if err != nil {
		return err
	}

	return service.repository.Delete(*ctxu.ID, id)
}
