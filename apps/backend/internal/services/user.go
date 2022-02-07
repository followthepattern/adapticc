package services

import (
	"backend/internal/container"
	"backend/internal/models"
	repositories "backend/internal/repositories/database"
	"backend/internal/utils"
	"fmt"
)

type User struct {
	userRepository repositories.User
}

func ResolveUser(cont container.IContainer) (dependency *User, err error) {
	key := utils.GetKey(dependency)
	obj, err := cont.Resolve(key)
	if err != nil {
		return nil, fmt.Errorf("can't resolve %T, error: %s", dependency, err.Error())
	}

	dependency, ok := obj.(*User)
	if !ok {
		return nil, fmt.Errorf("%T can't be resolved to %T", obj, dependency)
	}

	return dependency, nil
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
