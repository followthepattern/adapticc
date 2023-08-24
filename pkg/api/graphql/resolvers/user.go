package resolvers

import (
	"context"
	"net/http"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/models"
)

type User struct {
	ID        *string
	Email     *string
	FirstName *string
	LastName  *string
	Active    *bool
	UserLog
}

func getFromUserListResponseModel(response models.UserListResponse) *ListResponse[models.User] {
	resp := fromListReponseModel[models.User, models.User](models.ListResponse[models.User](response))
	resp.Data = response.Data
	return &resp
}

type UserResolver struct {
	cont *container.Container
	ctrl *controllers.User
}

func NewUserQuery(cont *container.Container) (*UserResolver, error) {
	ctrl, err := container.Resolve[controllers.User](cont)
	if err != nil {
		return nil, err
	}

	return &UserResolver{cont: cont, ctrl: ctrl}, nil
}

func (resolver UserResolver) Single(ctx context.Context, args struct{ Id string }) (*models.User, error) {
	return resolver.ctrl.GetByID(ctx, args.Id)
}

func (resolver UserResolver) List(ctx context.Context, args struct {
	Pagination *Pagination
	Filter     *models.ListFilter
	OrderBy    *[]models.OrderBy
}) (*ListResponse[models.User], error) {
	request := models.UserListRequestParams{}

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

	users, err := resolver.ctrl.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	result := getFromUserListResponseModel(users)

	return result, err
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
		Code: NewUint(http.StatusOK),
	}, nil
}

func (resolver UserResolver) Update(ctx context.Context, args struct {
	Id    string
	Model models.User
}) (*ResponseStatus, error) {
	user := models.User{
		ID:        &args.Id,
		FirstName: args.Model.FirstName,
		LastName:  args.Model.LastName,
	}

	err := resolver.ctrl.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return &ResponseStatus{
		Code: NewUint(200),
	}, nil
}

func (resolver UserResolver) Delete(ctx context.Context, args struct {
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
