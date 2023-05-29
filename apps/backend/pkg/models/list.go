package models

type ListFilter struct {
	PageSize *uint
	Page     *uint
}

type ListResponse[T any] struct {
	Count    int64 `json:"count"`
	PageSize *uint `json:"page_size"`
	Page     *uint `json:"page"`
	Data     []T   `json:"data"`
}
