package models

import "github.com/followthepattern/adapticc/types"

const (
	DefaultPageSize = 20
	DefaultPage     = 1
)

type ListRequestParams[T any] struct {
	Filter     T
	Pagination Pagination
	OrderBy    []OrderBy
}

type ListFilter struct {
	Search types.String
}

type Pagination struct {
	PageSize types.Uint
	Page     types.Uint
}

func (p *Pagination) SetDefaultIfEmpty() {
	if p.PageSize.Data < 1 {
		p.PageSize = types.UintFrom(DefaultPageSize)
	}

	if p.Page.Data < 1 {
		p.Page = types.UintFrom(DefaultPage)
	}
}

type OrderBy struct {
	Name string
	Desc types.Bool
}

type ListResponse[T any] struct {
	Count    types.Int64 `json:"count"`
	PageSize types.Uint  `json:"page_size"`
	Page     types.Uint  `json:"page"`
	Data     []T         `json:"data"`
}
