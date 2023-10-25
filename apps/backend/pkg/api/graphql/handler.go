package graphql

import (
	"net/http"

	"github.com/followthepattern/adapticc/pkg/api/graphql/resolvers"
	"github.com/followthepattern/adapticc/pkg/controllers"

	"github.com/followthepattern/graphql-go"
	"github.com/followthepattern/graphql-go/relay"
)

func New(controllers controllers.Controllers) http.Handler {
	resolverConfig := resolvers.NewResolverConfig(
		controllers.User(),
		controllers.Auth(),
		controllers.Product(),
		controllers.Role())

	resolver := resolvers.New(resolverConfig)

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}

	schema := graphql.MustParseSchema(Schema, &resolver, opts...)
	return &relay.Handler{Schema: schema}
}
