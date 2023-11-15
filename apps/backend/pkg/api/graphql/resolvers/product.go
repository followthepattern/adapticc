package resolvers

import (
	"context"
	"net/http"

	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/types"
)

type ProductResolver struct {
	ctrl controllers.Product
}

func NewProductQuery(ctrl controllers.Product) ProductResolver {
	return ProductResolver{ctrl: ctrl}
}

func (resolver ProductResolver) Single(ctx context.Context, args struct{ Id string }) (*models.Product, error) {
	product, err := resolver.ctrl.GetByID(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (resolver ProductResolver) List(ctx context.Context, args struct {
	Pagination *models.Pagination
	Filter     *models.ListFilter
	OrderBy    *[]models.OrderBy
}) (*models.ListResponse[models.Product], error) {
	request := models.ProductListRequestParams{}

	if args.Pagination != nil {
		request.Pagination = models.Pagination{
			PageSize: args.Pagination.PageSize,
			Page:     args.Pagination.Page,
		}
	}

	if args.Filter != nil {
		request.Filter = *args.Filter
	}

	if args.OrderBy != nil {
		request.OrderBy = *args.OrderBy
	}

	products, err := resolver.ctrl.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	return products, err
}

func (resolver ProductResolver) Create(ctx context.Context, args struct {
	Model models.Product
}) (*ResponseStatus, error) {
	err := resolver.ctrl.Create(ctx, args.Model)
	if err != nil {
		return nil, err
	}
	return &ResponseStatus{
		Code: types.IntFrom(http.StatusCreated),
	}, nil
}

func (resolver ProductResolver) Update(ctx context.Context, args struct {
	Id    types.String
	Model models.Product
}) (*ResponseStatus, error) {
	args.Model.ID = args.Id
	err := resolver.ctrl.Update(ctx, args.Model)
	if err != nil {
		return nil, err
	}
	return &ResponseStatus{
		Code: types.IntFrom(http.StatusOK),
	}, nil
}

func (resolver ProductResolver) Delete(ctx context.Context, args struct {
	Id string
}) (*ResponseStatus, error) {
	err := resolver.ctrl.Delete(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	return &ResponseStatus{
		Code: types.IntFrom(http.StatusOK),
	}, nil
}
