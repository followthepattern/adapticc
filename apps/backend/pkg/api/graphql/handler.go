package graphql

import (
	"net/http"

	"github.com/followthepattern/adapticc/pkg/api/graphql/resolvers"
	"github.com/followthepattern/adapticc/pkg/controllers"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func New(controllers controllers.Controllers) http.Handler {
	resolver := resolvers.New(controllers)

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}

	schema := graphql.MustParseSchema(Schema, &resolver, opts...)
	return &relay.Handler{Schema: schema}
}
