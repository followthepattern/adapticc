package models

const (
	DefaultPageSize = 20
	DefaultPage     = 1
)

type ListFilter struct {
	Search   *string
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
