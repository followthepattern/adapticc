package models

import "github.com/followthepattern/adapticc/pkg/types"

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
	Search string
}

type Pagination struct {
	PageSize types.NullUint
	Page     types.NullUint
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
	Desc *bool
}

type ListResponse[T any] struct {
	Count    types.NullInt64 `json:"count"`
	PageSize types.NullUint  `json:"page_size"`
	Page     types.NullUint  `json:"page"`
	Data     []T             `json:"data"`
}
