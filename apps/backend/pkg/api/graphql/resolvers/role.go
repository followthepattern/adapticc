package resolvers

import (
	"context"

	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/models"
)

func getFromRoleListResponseModel(response models.RoleListResponse) ListResponse[models.Role] {
	resp := fromListReponseModel[models.Role, models.Role](models.ListResponse[models.Role](response))
	resp.Data = response.Data
	return resp
}

type RoleResolver struct {
	ctrl controllers.Role
}

func NewRoleQuery(ctrl controllers.Role) RoleResolver {
	return RoleResolver{ctrl: ctrl}
}

func (resolver RoleResolver) Single(ctx context.Context, args struct{ Id string }) (*models.Role, error) {
	value, err := resolver.ctrl.GetByID(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (resolver RoleResolver) List(ctx context.Context, args struct {
	Pagination *Pagination
	Filter     *models.ListFilter
	OrderBy    *[]models.OrderBy
}) (*ListResponse[models.Role], error) {
	request := models.RoleListRequestParams{}

	if args.Pagination != nil {
		request.Pagination = models.Pagination{
			PageSize: args.Pagination.PageSize.ValuePtr(),
			Page:     args.Pagination.Page.ValuePtr(),
		}
	}

	if args.Filter != nil {
		request.Filter = *args.Filter
	}

	if args.OrderBy != nil {
		request.OrderBy = *args.OrderBy
	}

	values, err := resolver.ctrl.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	response := getFromRoleListResponseModel(*values)

	return &response, err
}
