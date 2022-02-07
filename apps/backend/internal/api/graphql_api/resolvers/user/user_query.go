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
	return GetFromModel(result), nil
}
