package graphql

import (
	"net/http"

	"github.com/followthepattern/adapticc/controllers"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func NewHandler(controllers controllers.Controllers, schemaDef string) http.Handler {
	resolver := NewResolver(controllers)

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}

	schema := graphql.MustParseSchema(schemaDef, &resolver, opts...)
	return &relay.Handler{Schema: schema}
}
