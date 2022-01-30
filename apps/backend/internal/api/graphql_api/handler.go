package graphql_api

import (
	"backend/internal/api/graphql_api/resolvers"
	"backend/internal/container"
	"backend/internal/utils"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func NewHandler(cont container.IContainer) (http.Handler, error) {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	resolverType := (*resolvers.Resolver)(nil)
	obj, err := cont.Resolve(utils.GetKey(resolverType))

	if err != nil {
		return nil, err
	}

	resolver := *obj.(*resolvers.Resolver)

	schema := graphql.MustParseSchema(Schema, &resolver, opts...)
	return &relay.Handler{Schema: schema}, nil
}
