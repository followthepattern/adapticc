package test_integration

var graphQLURL = "http://backend:8080/graphql"

type graphqlRequest struct {
	Query string `json:"query"`
}
