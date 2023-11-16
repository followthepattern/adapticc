package resolvers

import (
	"github.com/followthepattern/adapticc/pkg/controllers"
)

type Resolver struct {
	users        UserResolver
	products     ProductResolver
	roles        RoleResolver
	authMutation Auth
}

func New(controllers controllers.Controllers) Resolver {
	uq := NewUserQuery(controllers.User())
	am := NewAuthMutation(controllers.Auth())
	pq := NewProductQuery(controllers.Product())
	rq := NewRoleQuery(controllers.Role())

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

func (r Resolver) Authentication() (Auth, error) {
	return r.authMutation, nil
}

func (r Resolver) Products() (ProductResolver, error) {
	return r.products, nil
}

func (r Resolver) Roles() (RoleResolver, error) {
	return r.roles, nil
}
