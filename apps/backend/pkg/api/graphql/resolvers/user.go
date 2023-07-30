package resolvers

import (
	"context"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/graph-gophers/graphql-go"
)

type User struct {
	ID           *string
	Email        *string
	FirstName    *string
	LastName     *string
	Active       *bool
	RegisteredAt *graphql.Time
}

func getFromModel(model *models.User) *User {
	if model == nil || model.IsNil() {
		return nil
	}

	result := User{
		ID:           model.ID,
		Email:        model.Email,
		FirstName:    model.FirstName,
		LastName:     model.LastName,
		Active:       model.Active,
		RegisteredAt: utils.TimeToGraphqlTime(model.RegisteredAt),
	}

	return &result
}

func getFromModels(ms []models.User) []*User {
	result := make([]*User, len(ms))
	for i := 0; i < len(ms); i++ {
		result[i] = getFromModel(&ms[i])
	}
	return result
}

func getFromUserListResponseModel(response models.UserListResponse) ListResponse[*User] {
	resp := fromListReponseModel[models.User, *User](models.ListResponse[models.User](response))
	resp.Data = getFromModels(response.Data)
	return resp
}

type userListFilter struct {
	ListRequest
	Search *string
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

func (resolver UserResolver) Single(ctx context.Context, args struct{ Id string }) (*User, error) {
	u, err := resolver.ctrl.GetByID(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	user := getFromModel(u)
	return user, nil
}

func (resolver UserResolver) List(ctx context.Context, args struct{ Filter userListFilter }) (*ListResponse[*User], error) {
	filter := models.UserListRequestBody{
		ListFilter: models.ListFilter{
			PageSize: args.Filter.PageSize.ValuePtr(),
			Page:     args.Filter.Page.ValuePtr(),
			Search:   args.Filter.Search,
		},
	}

	users, err := resolver.ctrl.Get(ctx, filter)
	if err != nil {
		return nil, err
	}

	result := getFromUserListResponseModel(*users)

	return &result, err
}

func (resolver UserResolver) Profile(ctx context.Context) (*User, error) {
	u, err := resolver.ctrl.Profile(ctx)
	if err != nil {
		return nil, err
	}
	user := getFromModel(u)
	return user, nil
}

func (resolver UserResolver) Update(ctx context.Context, args struct {
	Id        string
	FirstName *string
	LastName  *string
}) (*ResponseStatus, error) {
	user := models.User{
		ID:        &args.Id,
		FirstName: args.FirstName,
		LastName:  args.LastName,
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
