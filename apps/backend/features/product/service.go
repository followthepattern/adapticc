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
	authorizationService auth.AuthorizationService
	logger               *slog.Logger
	productRepository    ProductDatabase
}

func NewProductService(cont container.Container, authorizationService auth.AuthorizationService) ProductService {
	productRepository := NewProductDatabase(cont.GetDB())
	product := ProductService{
		authorizationService: authorizationService,
		productRepository:    productRepository,
		logger:               cont.GetLogger(),
	}

	return product
}

func (service ProductService) GetByID(ctx context.Context, id string) (*ProductModel, error) {
	err := service.authorizationService.Authorize(ctx, accesscontrol.READ, id)
	if err != nil {
		return nil, err
	}

	return service.productRepository.GetByID(id)
}

func (service ProductService) Get(ctx context.Context, request ProductListRequestParams) (*ProductListResponse, error) {
	err := service.authorizationService.Authorize(ctx, accesscontrol.READ, accesscontrol.ALLRESOURCE)
	if err != nil {
		return nil, err
	}

	request.Pagination.SetDefaultIfEmpty()

	return service.productRepository.Get(request)
}

func (service ProductService) Create(ctx context.Context, value ProductModel) error {
	userID, err := service.authorizationService.AuthorizedUser(ctx, accesscontrol.CREATE, accesscontrol.NEW)
	if err != nil {
		return err
	}

	value.ID = types.StringFrom(uuid.NewString())
	value.Userlog.CreationUserID = types.StringFrom(userID)

	return service.productRepository.Create([]ProductModel{value})
}

func (service ProductService) Update(ctx context.Context, value ProductModel) error {
	userID, err := service.authorizationService.AuthorizedUser(ctx, accesscontrol.UPDATE, value.ID.Data)
	if err != nil {
		return err
	}

	value.UpdateUserID = types.StringFrom(userID)

	return service.productRepository.Update(value)
}

func (service ProductService) Delete(ctx context.Context, id string) error {
	err := service.authorizationService.Authorize(ctx, accesscontrol.DELETE, id)
	if err != nil {
		return err
	}

	return service.productRepository.Delete(id)
}
