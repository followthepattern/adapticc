package models

type ListResponse struct {
	Count    int64 `json:"count"`
	PageSize *uint `json:"page_size"`
	Page     *uint `json:"page"`
}
