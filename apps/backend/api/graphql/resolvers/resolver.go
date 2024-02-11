package resolvers

import (
	"github.com/followthepattern/adapticc/controllers"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/features/user"
)

type Resolver struct {
	users        user.UserResolver
	products     ProductResolver
	roles        RoleResolver
	authMutation auth.AuthResolver
}

func New(controllers controllers.Controllers) Resolver {
	uq := user.NewUserQuery(controllers.User())
	am := auth.NewAuthGraphQL(controllers.Auth())
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

func (r Resolver) Users() (user.UserResolver, error) {
	return r.users, nil
}

func (r Resolver) Authentication() (auth.AuthResolver, error) {
	return r.authMutation, nil
}

func (r Resolver) Products() (ProductResolver, error) {
	return r.products, nil
}

func (r Resolver) Roles() (RoleResolver, error) {
	return r.roles, nil
}
