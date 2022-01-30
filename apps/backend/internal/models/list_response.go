package models

type ListResponse struct {
	Count    int64       `json:"count"`
	PageSize int         `json:"page_size"`
	Page     int         `json:"page"`
	Data     interface{} `json:"data"`
}
