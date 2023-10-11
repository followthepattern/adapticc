package services

import (
	"context"
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
	Service
	productRepository database.Product
	roleRepository    database.Role
}

func NewProduct(ctx context.Context, cerbosClient cerbos.Client, productRepository database.Product, roleRepository database.Role, cfg config.Config, logger *zap.Logger) Product {
	base := Service{
		name:         "product",
		logger:       logger,
		cfg:          cfg,
		cerbosClient: cerbosClient,
	}

	product := Product{
		Service:           base,
		productRepository: productRepository,
		roleRepository:    roleRepository,
	}

	product.getRoles = product

	return product
}

func (service Product) GetByID(ctx context.Context, id string) (*models.Product, error) {
	ctxu, err := utils.GetUserContext(ctx)
	if err == nil {
		return nil, err
	}

	err = service.Authorize(ctx, *ctxu.ID, accesscontrol.READ, id)
	if err != nil {
		return nil, err
	}

	result, err := service.productRepository.GetByID(*ctxu.ID)
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

	err := service.Authorize(ctx, *ctxu.ID, accesscontrol.READ, accesscontrol.ALLRESOURCE)
	if err != nil {
		return nil, err
	}

	return service.productRepository.Get(filter)
}

func (service Product) Create(ctx context.Context, value models.Product) error {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return nil
	}

	value.ID = pointers.ToPtr(uuid.New().String())

	err = service.Authorize(ctx, *ctxu.ID, accesscontrol.CREATE, *value.ID)
	if err != nil {
		return err
	}

	return service.productRepository.Create(*ctxu.ID, []models.Product{value})
}

func (service Product) Update(ctx context.Context, value models.Product) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	err := service.Authorize(ctx, *ctxu.ID, accesscontrol.UPDATE, *value.ID)
	if err != nil {
		return err
	}

	return service.productRepository.Update(*ctxu.ID, value)
}

func (service Product) Delete(ctx context.Context, id string) error {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return nil
	}

	err = service.Authorize(ctx, *ctxu.ID, accesscontrol.DELETE, id)
	if err != nil {
		return err
	}

	return service.productRepository.Delete(id)
}

func (service Product) GetRolesByUserID(userID string) ([]string, error) {
	rolesArray, err := service.Service.GetRolesByUserID(userID)
	if err != nil {
		return nil, err
	}

	roles, err := service.roleRepository.GetRolesByUserID(userID)
	if err != nil {
		return nil, err
	}

	rolesArrayLen := len(rolesArray)
	result := make([]string, rolesArrayLen+len(roles))

	copy(result, rolesArray)

	for i, role := range roles {
		result[i+rolesArrayLen] = role.Name
	}

	return result, nil
}
