package internal

import (
	"backend/internal/api/graphql_api/resolvers"
	"backend/internal/container"
	"backend/internal/controllers"
	repositories "backend/internal/repositories/database"
	"backend/internal/services"
	"backend/internal/utils"
)

func RegisterDependencies(cont container.IContainer) (err error) {
	err = cont.Register(utils.GetKey((*repositories.User)(nil)), repositories.UserDependencyConstructor)
	if err != nil {
		return err
	}

	err = registerAuthentication(cont)
	if err != nil {
		return err
	}

	err = cont.Register(utils.GetKey((*services.User)(nil)), services.UserDependencyConstructor)
	if err != nil {
		return err
	}

	err = cont.Register(utils.GetKey((*controllers.UserController)(nil)), controllers.UserDependencyConstructor)
	if err != nil {
		return err
	}

	err = cont.Register(utils.GetKey((*resolvers.Resolver)(nil)), resolvers.ResolverDependencyConstructor)
	if err != nil {
		return err
	}

	return nil
}

func registerAuthentication(cont container.IContainer) (err error) {
	err = cont.Register(utils.GetKey((*repositories.UserToken)(nil)), repositories.UserTokenDependencyConstructor)
	if err != nil {
		return err
	}

	err = cont.Register(utils.GetKey((*services.Auth)(nil)), services.AuthDependencyConstructor)
	if err != nil {
		return err
	}

	err = cont.Register(utils.GetKey((*controllers.AuthController)(nil)), controllers.AuthDependencyConstructor)
	if err != nil {
		return err
	}

	return nil
}
