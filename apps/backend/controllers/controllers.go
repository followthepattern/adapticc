package controllers

import (
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/user"
)

type Controllers struct {
	user    user.UserController
	auth    Auth
	product Product
	role    Role
}

func New(cont container.Container) Controllers {
	return Controllers{
		user:    user.NewUserController(cont),
		auth:    NewAuth(cont),
		product: NewProduct(cont),
		role:    NewRole(cont),
	}
}

func (c Controllers) User() user.UserController {
	return c.user
}

func (c Controllers) Auth() Auth {
	return c.auth
}

func (c Controllers) Product() Product {
	return c.product
}

func (c Controllers) Role() Role {
	return c.role
}
