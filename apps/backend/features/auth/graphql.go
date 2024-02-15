package auth

import (
	"context"

	"github.com/followthepattern/adapticc/types"
	"github.com/graph-gophers/graphql-go"
)

type loginResponse struct {
	JWT       string       `json:"jwt,omitempty"`
	ExpiresAt graphql.Time `json:"expires_at,omitempty"`
}

func getFromLoginResponseModel(m LoginResponse) *loginResponse {
	expiresAt := graphql.Time{
		Time: m.ExpiresAt,
	}
	return &loginResponse{
		ExpiresAt: expiresAt,
		JWT:       m.JWT,
	}
}

type AuthGraphQL struct {
	auth AuthController
}

func NewAuthGraphQL(ctrl AuthController) AuthGraphQL {
	return AuthGraphQL{
		auth: ctrl,
	}
}

func (r AuthGraphQL) Login(ctx context.Context, args struct {
	Email    types.String
	Password types.String
}) (*loginResponse, error) {
	loginRequest := LoginRequestParams{
		Email:    args.Email,
		Password: args.Password,
	}

	loginResponse, err := r.auth.Login(ctx, loginRequest)
	if err != nil {
		return nil, err
	}

	return getFromLoginResponseModel(*loginResponse), nil
}

func (r AuthGraphQL) Register(ctx context.Context, args struct {
	Email     types.String
	FirstName types.String
	LastName  types.String
	Password  types.String
}) (*RegisterResponse, error) {
	registerRequest := RegisterRequestParams{
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
