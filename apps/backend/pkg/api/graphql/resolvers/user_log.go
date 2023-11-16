package resolvers

import "github.com/graph-gophers/graphql-go"

type UserLog struct {
	CreationUserID *string
	UpdateUserID   *string
	CreatedAt      *graphql.Time
	UpdatedAt      *graphql.Time
}
