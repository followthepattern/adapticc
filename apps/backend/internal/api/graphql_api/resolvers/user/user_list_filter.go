package user

import (
	"backend/internal/api/graphql_api/resolvers/base"
	"backend/internal/models"
)

type userListFilter struct {
	base.ListRequest
	Search *string
}

func getModelFromUserListFilter(request userListFilter) models.UserListRequest {
	return models.UserListRequest{
		ListRequest: base.GetModelFromListRequest(request.ListRequest),
		Search:      request.Search,
	}
}
