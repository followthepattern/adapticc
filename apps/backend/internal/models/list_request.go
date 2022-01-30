package models

import "github.com/doug-martin/goqu/v9"

type ListRequest struct {
	Filter   goqu.Ex `json:"filter"`
	PageSize *int    `json:"page_size"`
	Page     *int    `json:"page"`
}

func (lr ListRequest) GetFilter() goqu.Expression {
	return lr.Filter
}

func (lr *ListRequest) AddFilter(name string, value interface{}) {
	lr.Filter[name] = value
}
