package base

import (
	"backend/internal/api/graphql_api/resolvers/base/utils"
	"backend/internal/models"
)

type ListRequest struct {
	PageSize *utils.Uint
	Page     *utils.Uint
}

func GetModelFromListRequest(m ListRequest) models.ListRequest {
	lr := models.ListRequest{}

	if m.PageSize != nil {
		pageSize := m.PageSize.Value()
		lr.PageSize = &pageSize
	}

	if m.Page != nil {
		page := m.Page.Value()
		lr.Page = &page
	}

	return lr
}
