package resolvers

import "github.com/followthepattern/adapticc/pkg/models"

type ResponseStatus struct {
	Code Uint `json:"code"`
}

type ListResponse[T any] struct {
	Count    Int64
	PageSize *Uint
	Page     *Uint
	Data     []T
}

func fromListReponseModel[In any, Out any](m models.ListResponse[In]) ListResponse[Out] {
	lr := ListResponse[Out]{
		Count: NewInt64(m.Count),
	}

	if m.PageSize != nil {
		pageSize := NewUint(*m.PageSize)
		lr.PageSize = &pageSize
	}

	if m.Page != nil {
		page := NewUint(*m.Page)
		lr.Page = &page
	}

	return lr
}
