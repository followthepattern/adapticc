package product

import (
	"context"
	"log/slog"

	"github.com/followthepattern/adapticc/accesscontrol"
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/types"
	"github.com/google/uuid"
)

type ProductService struct {
	ac                accesscontrol.AccessControl
	logger            *slog.Logger
	productRepository ProductDatabase
}

func NewProductService(cont container.Container, productRepository ProductDatabase) ProductService {
	product := ProductService{
		ac:                cont.GetAccessControl().WithKind("product"),
		productRepository: productRepository,
		logger:            cont.GetLogger(),
	}

	return product
}

func (service ProductService) GetByID(ctx context.Context, id string) (*ProductModel, error) {
	ctxu, err := auth.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	err = service.ac.Authorize(ctx, ctxu.ID.Data, accesscontrol.READ, id, roles...)
	if err != nil {
		return nil, err
	}

	return service.productRepository.GetByID(id)
}

func (service ProductService) Get(ctx context.Context, request ProductListRequestParams) (*ProductListResponse, error) {
	ctxu, err := auth.GetUserContext(ctx)
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

func (service ProductService) Create(ctx context.Context, value ProductModel) error {
	ctxu, err := auth.GetUserContext(ctx)
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

	return service.productRepository.Create([]ProductModel{value})
}

func (service ProductService) Update(ctx context.Context, value ProductModel) error {
	ctxu, err := auth.GetUserContext(ctx)
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

func (service ProductService) Delete(ctx context.Context, id string) error {
	ctxu, err := auth.GetUserContext(ctx)
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
