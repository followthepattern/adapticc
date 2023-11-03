package services

import (
	"context"
	"log/slog"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/google/uuid"
)

type Product struct {
	ac                accesscontrol.AccessControl
	logger            *slog.Logger
	productRepository database.Product
	roleRepository    database.Role
}

func NewProduct(cont container.Container, productRepository database.Product, roleRepository database.Role) Product {
	product := Product{
		ac:                cont.GetAccessControl().WithKind("product"),
		productRepository: productRepository,
		roleRepository:    roleRepository,
		logger:            cont.GetLogger(),
	}

	return product
}

func (service Product) GetByID(ctx context.Context, id string) (*models.Product, error) {
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

	result, err := service.productRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if result.IsDefault() {
		return nil, nil
	}

	return result, nil
}

func (service Product) Get(ctx context.Context, filter models.ProductListRequestParams) (*models.ProductListResponse, error) {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return nil, err
	}

	err = service.ac.Authorize(ctx, ctxu.ID, accesscontrol.READ, accesscontrol.ALLRESOURCE, roles...)
	if err != nil {
		return nil, err
	}

	return service.productRepository.Get(filter)
}

func (service Product) Create(ctx context.Context, value models.Product) error {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return err
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
	value.Userlog.CreationUserID = ctxu.ID

	return service.productRepository.Create([]models.Product{value})
}

func (service Product) Update(ctx context.Context, value models.Product) error {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return err
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

	return service.productRepository.Update(value)
}

func (service Product) Delete(ctx context.Context, id string) error {
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

	return service.productRepository.Delete(id)
}
