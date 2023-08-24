package models

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
	Search *string
}

type Pagination struct {
	PageSize *uint
	Page     *uint
}

type OrderBy struct {
	Name string
	Desc *bool
}

type ListResponse[T any] struct {
	Count    int64 `json:"count"`
	PageSize *uint `json:"page_size"`
	Page     *uint `json:"page"`
	Data     []T   `json:"data"`
}
