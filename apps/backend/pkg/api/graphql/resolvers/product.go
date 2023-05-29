package resolvers

import (
	"context"
	"net/http"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/models"
)

type ProductListFilter struct {
	ListRequest
	ProductID *string
}

func getFromProductListResponseModel(response models.ProductListResponse) ListResponse[models.Product] {
	resp := fromListReponseModel[models.Product, models.Product](models.ListResponse[models.Product](response))
	resp.Data = response.Data
	return resp
}

type ProductResolver struct {
	cont *container.Container
	ctrl *controllers.Product
}

func NewProductQuery(cont *container.Container) (*ProductResolver, error) {
	ctrl, err := container.Resolve[controllers.Product](cont)
	if err != nil {
		return nil, err
	}

	return &ProductResolver{cont: cont, ctrl: ctrl}, nil
}

func (resolver ProductResolver) Single(ctx context.Context, args struct{ ProductID string }) (*models.Product, error) {
	product, err := resolver.ctrl.GetByID(ctx, args.ProductID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (resolver ProductResolver) List(ctx context.Context, args struct{ Filter ProductListFilter }) (*ListResponse[models.Product], error) {
	filter := models.ProductListRequestBody{
		ListFilter: models.ListFilter{
			PageSize: &args.Filter.PageSize.uint,
			Page:     &args.Filter.Page.uint,
		},
		ProductRequestBody: models.ProductRequestBody{
			ProductID: args.Filter.ProductID,
		},
	}

	products, err := resolver.ctrl.Get(ctx, filter)
	if err != nil {
		return nil, err
	}

	response := getFromProductListResponseModel(*products)

	return &response, err
}

func (resolver ProductResolver) Create(ctx context.Context, args struct {
	Title       *string
	Description *string
}) (*ResponseStatus, error) {
	Product := models.Product{
		Title:       args.Title,
		Description: args.Description,
	}

	err := resolver.ctrl.Create(ctx, Product)
	if err != nil {
		return nil, err
	}
	return &ResponseStatus{
		Code: NewUint(http.StatusOK),
	}, nil
}

func (resolver ProductResolver) Update(ctx context.Context, args struct {
	ProductID   string
	Title       *string
	Description *string
}) (*ResponseStatus, error) {
	Product := models.Product{
		ProductID:   &args.ProductID,
		Title:       args.Title,
		Description: args.Description,
	}

	err := resolver.ctrl.Update(ctx, Product)
	if err != nil {
		return nil, err
	}
	return &ResponseStatus{
		Code: NewUint(200),
	}, nil
}

func (resolver ProductResolver) Delete(ctx context.Context, args struct {
	ProductID string
}) (*ResponseStatus, error) {
	err := resolver.ctrl.Delete(ctx, args.ProductID)
	if err != nil {
		return nil, err
	}
	return &ResponseStatus{
		Code: NewUint(200),
	}, nil
}
