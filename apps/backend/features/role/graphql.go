package role

import (
	"context"
	"net/http"

	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
)

type RoleResolver struct {
	ctrl RoleController
}

func NewRoleQuery(ctrl RoleController) RoleResolver {
	return RoleResolver{ctrl: ctrl}
}

func (resolver RoleResolver) Single(ctx context.Context, args struct{ Id string }) (*RoleModel, error) {
	value, err := resolver.ctrl.GetByID(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (resolver RoleResolver) List(ctx context.Context, args struct {
	Pagination *models.Pagination
	Filter     *models.ListFilter
	OrderBy    *[]models.OrderBy
}) (*models.ListResponse[RoleModel], error) {
	request := RoleListRequestParams{}

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

	values, err := resolver.ctrl.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	return values, nil
}

func (resolver RoleResolver) AddRoleToUser(ctx context.Context, args struct {
	UserID types.String
	RoleID types.String
}) (*models.ResponseStatus, error) {
	err := resolver.ctrl.AddRoleToUser(ctx, UserRoleModel{UserID: args.UserID, RoleID: args.RoleID})
	if err != nil {
		return nil, err
	}
	return &models.ResponseStatus{
		Code: http.StatusCreated,
	}, nil
}

func (resolver RoleResolver) DeleteRoleFromUser(ctx context.Context, args struct {
	UserID types.String
	RoleID types.String
}) (*models.ResponseStatus, error) {
	err := resolver.ctrl.DeleteRoleFromUser(ctx, UserRoleModel{UserID: args.UserID, RoleID: args.RoleID})
	if err != nil {
		return nil, err
	}
	return &models.ResponseStatus{
		Code: http.StatusOK,
	}, nil
}
