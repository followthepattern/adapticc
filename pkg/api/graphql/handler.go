package graphql

import (
	"net/http"

	"github.com/followthepattern/adapticc/pkg/api/graphql/resolvers"

	"github.com/followthepattern/graphql-go"
	"github.com/followthepattern/graphql-go/relay"
)

func New(resolver *resolvers.Resolver) http.Handler {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}

	schema := graphql.MustParseSchema(Schema, resolver, opts...)
	return &relay.Handler{Schema: schema}
}
