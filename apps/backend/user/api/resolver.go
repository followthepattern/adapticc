package api

import (
	"context"
	"net/http"

	"github.com/followthepattern/adapticc/api/graphql/resolvers"
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
	"github.com/followthepattern/adapticc/user"
)

type UserResolver struct {
	ctrl user.UserController
}

func NewUserQuery(ctrl user.UserController) UserResolver {
	return UserResolver{ctrl: ctrl}
}

func (resolver UserResolver) Single(ctx context.Context, args struct{ Id types.String }) (*user.UserModel, error) {
	return resolver.ctrl.GetByID(ctx, args.Id)
}

func (resolver UserResolver) List(ctx context.Context, args struct {
	Pagination *models.Pagination
	Filter     *models.ListFilter
	OrderBy    *[]models.OrderBy
}) (*models.ListResponse[user.UserModel], error) {
	request := user.UserListRequestParams{}

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

	results, err := resolver.ctrl.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	return results, err
}

func (resolver UserResolver) Profile(ctx context.Context) (*user.UserModel, error) {
	return resolver.ctrl.Profile(ctx)
}

func (resolver UserResolver) Create(ctx context.Context, args struct {
	Model user.UserModel
}) (*resolvers.ResponseStatus, error) {
	err := resolver.ctrl.Create(ctx, args.Model)
	if err != nil {
		return nil, err
	}
	return &resolvers.ResponseStatus{
		Code: http.StatusCreated,
	}, nil
}

func (resolver UserResolver) Update(ctx context.Context, args struct {
	Id    types.String
	Model user.UserModel
}) (*resolvers.ResponseStatus, error) {
	args.Model.ID = args.Id
	err := resolver.ctrl.Update(ctx, args.Model)
	if err != nil {
		return nil, err
	}
	return &resolvers.ResponseStatus{
		Code: http.StatusOK,
	}, nil
}

func (resolver UserResolver) Delete(ctx context.Context, args struct {
	Id types.String
}) (*resolvers.ResponseStatus, error) {
	err := resolver.ctrl.Delete(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	return &resolvers.ResponseStatus{
		Code: http.StatusOK,
	}, nil
}
