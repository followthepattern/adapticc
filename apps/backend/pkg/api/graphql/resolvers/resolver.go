package resolvers

import (
	"errors"

	"github.com/followthepattern/adapticc/pkg/container"
)

type Resolver struct {
	users        UserResolver
	products     ProductResolver
	authMutation AuthMutation
}

func ResolverDependencyConstructor(cont *container.Container) (*Resolver, error) {
	uq, err := NewUserQuery(cont)
	if err != nil {
		return nil, err
	}
	if uq == nil {
		return nil, errors.New("userQuery can't be nil")
	}

	am, err := NewAuthMutation(cont)
	if err != nil {
		return nil, err
	}
	if am == nil {
		return nil, errors.New("authMutation can't be nil")
	}

	pq, err := NewProductQuery(cont)
	if err != nil {
		return nil, err
	}

	if pq == nil {
		return nil, errors.New("productQuery can't be nil")
	}

	resolver := Resolver{
		users:        *uq,
		products:     *pq,
		authMutation: *am,
	}

	return &resolver, nil
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
