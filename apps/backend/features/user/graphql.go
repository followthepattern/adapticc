package user

import (
	"context"
	"net/http"

	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
)

type UserGraphQL struct {
	ctrl UserController
}

func NewUserGraphQL(ctrl UserController) UserGraphQL {
	return UserGraphQL{ctrl: ctrl}
}

func (graphQL UserGraphQL) Single(ctx context.Context, args struct{ Id types.String }) (*UserModel, error) {
	return graphQL.ctrl.GetByID(ctx, args.Id)
}

func (graphQL UserGraphQL) List(ctx context.Context, args struct {
	Pagination *models.Pagination
	Filter     *models.ListFilter
	OrderBy    *[]models.OrderBy
}) (*models.ListResponse[UserModel], error) {
	request := UserListRequestParams{}

	if args.Pagination != nil {
		request.Pagination = models.Pagination{
			PageSize: args.Pagination.PageSize,
			Page:     args.Pagination.Page,
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

	results, err := graphQL.ctrl.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	return results, err
}

func (graphQL UserGraphQL) Profile(ctx context.Context) (*UserModel, error) {
	return graphQL.ctrl.Profile(ctx)
}

func (graphQL UserGraphQL) Create(ctx context.Context, args struct {
	Model UserModel
}) (*models.ResponseStatus, error) {
	err := graphQL.ctrl.Create(ctx, args.Model)
	if err != nil {
		return nil, err
	}
	return &models.ResponseStatus{
		Code: http.StatusCreated,
	}, nil
}

func (graphQL UserGraphQL) Update(ctx context.Context, args struct {
	Id    types.String
	Model UserModel
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

func (graphQL UserGraphQL) Delete(ctx context.Context, args struct {
	Id types.String
}) (*models.ResponseStatus, error) {
	err := graphQL.ctrl.Delete(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	return &models.ResponseStatus{
		Code: http.StatusOK,
	}, nil
}
