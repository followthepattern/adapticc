package auth

import (
	"backend/internal/controllers"
	"backend/internal/models"
	"context"
)

type AuthMutation struct {
	ac controllers.AuthController
}

func NewAuthMutation(ac controllers.AuthController) AuthMutation {
	return AuthMutation{
		ac: ac,
	}
}

func (am AuthMutation) Login(ctx context.Context, args struct {
	Email    string
	Password string
}) (*loginResponse, error) {
	result, err := am.ac.Login(args.Email, args.Password)
	if err != nil {
		return nil, err
	}
	return GetFromModel(*result), nil
}

func (am AuthMutation) Register(ctx context.Context, args struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}) (*models.RegisterResponse, error) {
	return am.ac.Register(args.Email, args.FirstName, args.LastName, args.Password)
}
