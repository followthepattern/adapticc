package base

import (
	"backend/internal/api/graphql_api/resolvers/base/utils"
	"backend/internal/models"
)

type ListResponse struct {
	Count    utils.Int64
	PageSize *utils.Uint
	Page     *utils.Uint
}

func FromListReponseModel(m models.ListResponse) ListResponse {
	lr := ListResponse{
		Count: utils.NewInt64(m.Count),
	}

	if m.PageSize != nil {
		pageSize := utils.NewUint(*m.PageSize)
		lr.PageSize = &pageSize
	}

	if m.Page != nil {
		page := utils.NewUint(*m.Page)
		lr.Page = &page
	}

	return lr
}
