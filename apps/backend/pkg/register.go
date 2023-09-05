package internal

import (
	"github.com/followthepattern/adapticc/pkg/api/graphql/resolvers"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/controllers"
	repositories "github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/services"
)

func RegisterDependencies(cont *container.Container) error {
	err := container.Register(cont, repositories.UserDependencyConstructor)
	if err != nil {
		return err
	}

	err = container.Register(cont, controllers.UserDependencyConstructor)
	if err != nil {
		return err
	}

	// auth plugin
	err = container.Register(cont, repositories.AuthDependencyConstructor)
	if err != nil {
		return err
	}

	err = container.Register(cont, services.AuthDependencyConstructor)
	if err != nil {
		return err
	}

	err = container.Register(cont, controllers.AuthDependencyConstructor)
	if err != nil {
		return err
	}

	// mail plugin
	// services.RegisterMailChannel(cont)

	// err = container.Register(cont, services.MailDependencyConstructor)
	// if err != nil {
	// 	return err
	// }

	// product plugin
	err = container.Register(cont, repositories.ProductDependencyConstructor)
	if err != nil {
		return err
	}

	err = container.Register(cont, controllers.ProductDependencyConstructor)
	if err != nil {
		return err
	}

	err = container.Register(cont, resolvers.ResolverDependencyConstructor)
	if err != nil {
		return err
	}

	return nil
}