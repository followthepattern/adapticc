package services

import (
	"backend/internal/container"
	"backend/internal/models"
	repositories "backend/internal/repositories/database"
	"backend/internal/utils"
)

type User struct {
	userRepository repositories.User
}

func UserDependencyConstructor(cont container.IContainer) (interface{}, error) {
	us := User{}
	key := utils.GetKey((*repositories.User)(nil))
	obj, err := cont.Resolve(key)

	if err != nil {
		return nil, err
	}

	us.userRepository = *obj.(*repositories.User)
	return &us, nil
}

func (service User) GetByID(id string) (*models.User, error) {
	return service.userRepository.GetByID(id)
}

func (service User) GetByToken(token string) (*models.User, error) {
	return service.userRepository.GetByToken(token)
}
