package resolvers

import (
	"context"
	"net/http"

	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/types"
)

type UserResolver struct {
	ctrl controllers.User
}

func NewUserQuery(ctrl controllers.User) UserResolver {
	return UserResolver{ctrl: ctrl}
}

func (resolver UserResolver) Single(ctx context.Context, args struct{ Id types.String }) (*models.User, error) {
	return resolver.ctrl.GetByID(ctx, args.Id)
}

func (resolver UserResolver) List(ctx context.Context, args struct {
	Pagination *models.Pagination
	Filter     *models.ListFilter
	OrderBy    *[]models.OrderBy
}) (*models.ListResponse[models.User], error) {
	request := models.UserListRequestParams{}

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

func (resolver UserResolver) Profile(ctx context.Context) (*models.User, error) {
	return resolver.ctrl.Profile(ctx)
}

func (resolver UserResolver) Create(ctx context.Context, args struct {
	Model models.User
}) (*ResponseStatus, error) {
	err := resolver.ctrl.Create(ctx, args.Model)
	if err != nil {
		return nil, err
	}
	return &ResponseStatus{
		Code: http.StatusCreated,
	}, nil
}

func (resolver UserResolver) Update(ctx context.Context, args struct {
	Id    types.String
	Model models.User
}) (*ResponseStatus, error) {
	args.Model.ID = args.Id
	err := resolver.ctrl.Update(ctx, args.Model)
	if err != nil {
		return nil, err
	}
	return &ResponseStatus{
		Code: http.StatusOK,
	}, nil
}

func (resolver UserResolver) Delete(ctx context.Context, args struct {
	Id types.String
}) (*ResponseStatus, error) {
	err := resolver.ctrl.Delete(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	return &ResponseStatus{
		Code: http.StatusOK,
	}, nil
}
