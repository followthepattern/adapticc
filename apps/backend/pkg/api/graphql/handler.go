package graphql

import (
	"net/http"

	"github.com/followthepattern/adapticc/pkg/api/graphql/resolvers"
	"github.com/followthepattern/adapticc/pkg/container"

	"github.com/followthepattern/graphql-go"
	"github.com/followthepattern/graphql-go/relay"
)

func NewHandler(cont *container.Container) (http.Handler, error) {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	var resolver *resolvers.Resolver
	resolver, err := container.Resolve[resolvers.Resolver](cont)
	if err != nil {
		return nil, err
	}

	schema := graphql.MustParseSchema(Schema, resolver, opts...)
	return &relay.Handler{Schema: schema}, nil
}
