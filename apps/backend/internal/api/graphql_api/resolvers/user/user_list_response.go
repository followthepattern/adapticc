package user

import (
	"backend/internal/api/graphql_api/resolvers/base"
	"backend/internal/models"
)

type userListResponse struct {
	base.ListResponse
	Data []*user `json:"data"`
}

func getFromUserListResponseModel(response *models.UserListResponse) *userListResponse {
	return &userListResponse{
		ListResponse: base.FromListReponseModel(response.ListResponse),
		Data:         getFromModels(response.Data),
	}
}
