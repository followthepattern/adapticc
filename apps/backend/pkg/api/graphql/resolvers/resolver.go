package resolvers

import (
	"github.com/followthepattern/adapticc/pkg/controllers"
)

type Resolver struct {
	users        UserResolver
	products     ProductResolver
	roles        RoleResolver
	authMutation AuthMutation
}

type ResolverConfig struct {
	userController    controllers.User
	authController    controllers.Auth
	productController controllers.Product
	roleController    controllers.Role
}

func NewResolverConfig(
	userController controllers.User,
	authController controllers.Auth,
	productController controllers.Product,
	roleController controllers.Role,
) ResolverConfig {
	return ResolverConfig{
		userController:    userController,
		authController:    authController,
		productController: productController,
		roleController:    roleController,
	}
}

func New(rc ResolverConfig) Resolver {
	uq := NewUserQuery(rc.userController)
	am := NewAuthMutation(rc.authController)
	pq := NewProductQuery(rc.productController)
	rq := NewRoleQuery(rc.roleController)

	resolver := Resolver{
		users:        uq,
		products:     pq,
		authMutation: am,
		roles:        rq,
	}

	return resolver
}

func (r Resolver) Users() (UserResolver, error) {
	return r.users, nil
}

func (r Resolver) Authentication() (AuthMutation, error) {
	return r.authMutation, nil
}

func (r Resolver) Products() (ProductResolver, error) {
	return r.products, nil
}

func (r Resolver) Roles() (RoleResolver, error) {
	return r.roles, nil
}
