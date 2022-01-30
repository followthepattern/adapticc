package controllers

import (
	"backend/internal/container"
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
)

type AuthController struct {
	auth services.Auth
}

func AuthDependencyConstructor(cont container.IContainer) (interface{}, error) {
	key := utils.GetKey((*services.Auth)(nil))
	obj, err := cont.Resolve(key)

	if err != nil {
		return nil, err
	}
	return &AuthController{auth: *obj.(*services.Auth)}, nil
}

func (ac AuthController) Login(email string, password string) (*models.LoginResponse, error) {
	return ac.auth.Login(email, password)
}

func (ac AuthController) Register(email string, firstName string, lastName string, password string) (*models.RegisterResponse, error) {
	return ac.auth.Register(email, firstName, lastName, password)
}
