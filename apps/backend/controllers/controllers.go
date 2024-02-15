package controllers

import (
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/features/product"
	"github.com/followthepattern/adapticc/features/role"
	"github.com/followthepattern/adapticc/features/user"
)

type Controllers struct {
	user    user.UserController
	auth    auth.AuthController
	product product.ProductController
	role    role.RoleController
}

func New(cont container.Container) Controllers {
	return Controllers{
		user:    user.NewUserController(cont),
		auth:    auth.NewAuthController(cont),
		product: product.NewProductController(cont),
		role:    role.NewRoleController(cont),
	}
}

func (c Controllers) User() user.UserController {
	return c.user
}

func (c Controllers) Auth() auth.AuthController {
	return c.auth
}

func (c Controllers) Product() product.ProductController {
	return c.product
}

func (c Controllers) Role() role.RoleController {
	return c.role
}
