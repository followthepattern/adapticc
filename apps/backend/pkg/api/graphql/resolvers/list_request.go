package resolvers

import (
	"github.com/followthepattern/adapticc/pkg/models"
)

type ListRequest struct {
	Search   *string
	PageSize *Uint
	Page     *Uint
}

func GetModelFromListRequest(m ListRequest) models.ListFilter {
	lr := models.ListFilter{}

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
