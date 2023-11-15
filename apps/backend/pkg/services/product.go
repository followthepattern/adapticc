package services

import (
	"context"
	"log/slog"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/types"
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

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.READ, id, roles...)
	if err != nil {
		return nil, err
	}

	return service.productRepository.GetByID(id)
}

func (service Product) Get(ctx context.Context, request models.ProductListRequestParams) (*models.ProductListResponse, error) {
	ctxu, err := utils.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	roles, err := service.roleRepository.GetRoleCodes(ctxu.ID)
	if err != nil {
		return nil, err
	}

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.READ, accesscontrol.ALLRESOURCE, roles...)
	if err != nil {
		return nil, err
	}

	request.Pagination.SetDefaultIfEmpty()

	return service.productRepository.Get(request)
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

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.CREATE, accesscontrol.NEW, roles...)
	if err != nil {
		return err
	}

	value.ID = types.StringFrom(uuid.NewString())
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

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.UPDATE, value.ID.Data, roles...)
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

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.DELETE, id, roles...)
	if err != nil {
		return err
	}

	return service.productRepository.Delete(id)
}
