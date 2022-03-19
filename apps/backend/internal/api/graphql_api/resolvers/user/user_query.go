package user

import (
	"backend/internal/controllers"
	"context"
)

type UserQuery struct {
	userController controllers.UserController
}

func NewUserQuery(uc controllers.UserController) UserQuery {
	return UserQuery{
		userController: uc,
	}
}

func (uq UserQuery) Single(ctx context.Context, args struct{ Id string }) (*user, error) {
	result, err := uq.userController.GetByID(args.Id)
	if err != nil {
		return nil, err
	}
	return getFromModel(result), nil
}

func (uq UserQuery) List(ctx context.Context, args struct{ Filter userListFilter }) (*userListResponse, error) {
	result, err := uq.userController.Get(getModelFromUserListFilter(args.Filter))
	if err != nil {
		return nil, err
	}

	return getFromUserListResponseModel(result), nil
}

func (uq UserQuery) Profile(ctx context.Context) (*user, error) {
	result, err := uq.userController.Profile(ctx)

	if err != nil {
		return nil, err
	}
	return getFromModel(result), nil
}
