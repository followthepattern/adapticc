package resolvers

import (
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
)

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
		lr.PageSize = pointers.ToPtr(NewUint(*m.PageSize))
	}

	if m.Page != nil {
		lr.Page = pointers.ToPtr(NewUint(*m.Page))
	}

	return lr
}
