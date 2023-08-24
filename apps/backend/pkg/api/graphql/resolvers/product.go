package resolvers

import (
	"context"
	"net/http"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
)

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

func (resolver ProductResolver) Single(ctx context.Context, args struct{ Id string }) (*models.Product, error) {
	product, err := resolver.ctrl.GetByID(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (resolver ProductResolver) List(ctx context.Context, args struct {
	Pagination *Pagination
	Filter     *models.ListFilter
	OrderBy    *[]models.OrderBy
}) (*ListResponse[models.Product], error) {
	request := models.ProductListRequestParams{}

	if args.Pagination != nil {
		request.Pagination = models.Pagination{
			PageSize: args.Pagination.PageSize.ValuePtr(),
			Page:     args.Pagination.Page.ValuePtr(),
		}
	}

	if args.Filter != nil {
		request.Filter = models.ListFilter{
			Search: args.Filter.Search,
		}
	}

	if args.OrderBy != nil {
		request.OrderBy = *args.OrderBy
	}

	products, err := resolver.ctrl.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	response := getFromProductListResponseModel(*products)

	return &response, err
}

func (resolver ProductResolver) Create(ctx context.Context, args struct {
	Model models.Product
}) (*ResponseStatus, error) {
	err := resolver.ctrl.Create(ctx, args.Model)
	if err != nil {
		return nil, err
	}
	return &ResponseStatus{
		Code: NewUint(http.StatusOK),
	}, nil
}

func (resolver ProductResolver) Update(ctx context.Context, args struct {
	Id    string
	Model models.Product
}) (*ResponseStatus, error) {
	args.Model.ID = pointers.ToPtr(args.Id)
	err := resolver.ctrl.Update(ctx, args.Model)
	if err != nil {
		return nil, err
	}
	return &ResponseStatus{
		Code: NewUint(200),
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
		Code: NewUint(200),
	}, nil
}
