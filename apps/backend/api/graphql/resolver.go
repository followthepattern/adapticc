package graphql

import (
	"github.com/followthepattern/adapticc/controllers"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/features/product"
	"github.com/followthepattern/adapticc/features/role"
	"github.com/followthepattern/adapticc/features/user"
)

type Resolver struct {
	users        user.UserGraphQL
	products     product.ProductGraphQL
	roles        role.RoleGraphQL
	authMutation auth.AuthGraphQL
}

func NewResolver(controllers controllers.Controllers) Resolver {
	uq := user.NewUserGraphQL(controllers.User())
	am := auth.NewAuthGraphQL(controllers.Auth())
	pq := product.NewProductGraphQL(controllers.Product())
	rq := role.NewroleGraphql(controllers.Role())

	resolver := Resolver{
		users:        uq,
		products:     pq,
		authMutation: am,
		roles:        rq,
	}

	return resolver
}

func (r Resolver) Users() (user.UserGraphQL, error) {
	return r.users, nil
}

func (r Resolver) Authentication() (auth.AuthGraphQL, error) {
	return r.authMutation, nil
}

func (r Resolver) Products() (product.ProductGraphQL, error) {
	return r.products, nil
}

func (r Resolver) Roles() (role.RoleGraphQL, error) {
	return r.roles, nil
}
