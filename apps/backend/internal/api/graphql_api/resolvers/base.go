package resolvers

import (
	"backend/internal/api/graphql_api/resolvers/auth"
	"backend/internal/api/graphql_api/resolvers/user"
	"backend/internal/container"
	"backend/internal/controllers"
	"backend/internal/utils"
	"fmt"
)

type Resolver struct {
	uc controllers.UserController
	ac controllers.AuthController
}

func ResolverDependencyConstructor(cont container.IContainer) (interface{}, error) {
	uc, err := resolveUserController(cont)
	if err != nil {
		return nil, err
	}

	ac, err := resolveAuthController(cont)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		uc: *uc,
		ac: *ac,
	}, nil
}

func resolveUserController(cont container.IContainer) (*controllers.UserController, error) {
	var dependency *controllers.UserController
	key := utils.GetKey(dependency)
	obj, err := cont.Resolve(key)

	if err != nil {
		return nil, err
	}

	if result, ok := obj.(*controllers.UserController); ok {
		return result, nil
	}

	return nil, fmt.Errorf("can't resolve %T", dependency)
}

func resolveAuthController(cont container.IContainer) (*controllers.AuthController, error) {
	dependency := (*controllers.AuthController)(nil)
	key := utils.GetKey(dependency)
	obj, err := cont.Resolve(key)

	if err != nil {
		return nil, err
	}

	if result, ok := obj.(*controllers.AuthController); ok {
		return result, nil
	}

	return nil, fmt.Errorf("can't resolve %T", dependency)
}

func (r Resolver) Users() (user.UserQuery, error) {
	return user.NewUserQuery(r.uc), nil
}

func (r Resolver) Authentication() (auth.AuthMutation, error) {
	return auth.NewAuthMutation(r.ac), nil
}
