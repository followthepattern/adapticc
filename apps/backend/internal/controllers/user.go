package controllers

import (
	"backend/internal/container"
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
)

type UserController struct {
	us services.User
}

func UserDependencyConstructor(cont container.IContainer) (interface{}, error) {
	key := utils.GetKey((*services.User)(nil))
	obj, err := cont.Resolve(key)

	if err != nil {
		return nil, err
	}
	return &UserController{us: *obj.(*services.User)}, nil
}

func (ctrl UserController) Create(user *models.User) error {
	return nil
}

func (ctrl UserController) GetByID(id string) (*models.User, error) {
	return ctrl.us.GetByID(id)
}

func (ctrl UserController) Get() ([]models.User, error) {
	return nil, nil
}

func (ctrl UserController) Update(user *models.User) error {
	return nil
}

func (ctrl UserController) Delete(id int) error {
	return nil
}
