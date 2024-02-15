package product

import (
	"context"
	"net/http"

	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
)

type ProductGraphQL struct {
	ctrl ProductController
}

func NewProductGraphQL(ctrl ProductController) ProductGraphQL {
	return ProductGraphQL{ctrl: ctrl}
}

func (graphQL ProductGraphQL) Single(ctx context.Context, args struct{ Id string }) (*ProductModel, error) {
	product, err := graphQL.ctrl.GetByID(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (graphQL ProductGraphQL) List(ctx context.Context, args struct {
	Pagination *models.Pagination
	Filter     *models.ListFilter
	OrderBy    *[]models.OrderBy
}) (*models.ListResponse[ProductModel], error) {
	request := ProductListRequestParams{}

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

	products, err := graphQL.ctrl.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	return products, err
}

func (graphQL ProductGraphQL) Create(ctx context.Context, args struct {
	Model ProductModel
}) (*models.ResponseStatus, error) {
	err := graphQL.ctrl.Create(ctx, args.Model)
	if err != nil {
		return nil, err
	}
	return &models.ResponseStatus{
		Code: http.StatusCreated,
	}, nil
}

func (graphQL ProductGraphQL) Update(ctx context.Context, args struct {
	Id    types.String
	Model ProductModel
}) (*models.ResponseStatus, error) {
	args.Model.ID = args.Id
	err := graphQL.ctrl.Update(ctx, args.Model)
	if err != nil {
		return nil, err
	}
	return &models.ResponseStatus{
		Code: http.StatusOK,
	}, nil
}

func (graphQL ProductGraphQL) Delete(ctx context.Context, args struct {
	Id string
}) (*models.ResponseStatus, error) {
	err := graphQL.ctrl.Delete(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	return &models.ResponseStatus{
		Code: http.StatusOK,
	}, nil
}
