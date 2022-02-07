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
	var userRepository *repositories.User
	err = cont.Register(utils.GetKey(userRepository), repositories.UserDependencyConstructor)
	if err != nil {
		return err
	}

	err = registerAuthentication(cont)
	if err != nil {
		return err
	}

	var userService *services.User
	err = cont.Register(utils.GetKey(userService), services.UserDependencyConstructor)
	if err != nil {
		return err
	}

	var userController *controllers.UserController
	err = cont.Register(utils.GetKey(userController), controllers.UserDependencyConstructor)
	if err != nil {
		return err
	}

	var resolver *resolvers.Resolver
	err = cont.Register(utils.GetKey(resolver), resolvers.ResolverDependencyConstructor)
	if err != nil {
		return err
	}

	return nil
}

func registerAuthentication(cont container.IContainer) (err error) {
	var userTokenRepository *repositories.UserToken
	err = cont.Register(utils.GetKey(userTokenRepository), repositories.UserTokenDependencyConstructor)
	if err != nil {
		return err
	}

	var authServices *services.Auth
	err = cont.Register(utils.GetKey(authServices), services.AuthDependencyConstructor)
	if err != nil {
		return err
	}

	var authController *controllers.AuthController
	err = cont.Register(utils.GetKey(authController), controllers.AuthDependencyConstructor)
	if err != nil {
		return err
	}

	return nil
}
