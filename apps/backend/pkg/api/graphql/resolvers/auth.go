package resolvers

import (
	"context"

	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/types"
)

type AuthMutation struct {
	auth controllers.Auth
}

func NewAuthMutation(ctrl controllers.Auth) AuthMutation {
	return AuthMutation{
		auth: ctrl,
	}
}

func (r AuthMutation) Login(ctx context.Context, args struct {
	Email    types.String
	Password types.String
}) (*loginResponse, error) {
	loginRequest := models.LoginRequestParams{
		Email:    args.Email,
		Password: args.Password,
	}

	loginResponse, err := r.auth.Login(ctx, loginRequest)
	if err != nil {
		return nil, err
	}

	return getFromLoginResponseModel(*loginResponse), nil
}

func (r AuthMutation) Register(ctx context.Context, args struct {
	Email     types.String
	FirstName types.String
	LastName  types.String
	Password  types.String
}) (*models.RegisterResponse, error) {
	registerRequest := models.RegisterRequestParams{
		Email:     args.Email,
		FirstName: args.FirstName,
		LastName:  args.LastName,
		Password:  args.Password,
	}

	registerResponse, err := r.auth.Register(ctx, registerRequest)
	if err != nil {
		return nil, err
	}

	return registerResponse, nil
}
