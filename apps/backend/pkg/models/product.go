package models

import (
	"github.com/followthepattern/adapticc/pkg/request"
)

type Product struct {
	ID          *string `db:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (m *Product) IsNil() bool {
	return m == nil || m.ID == nil
}

type ProductRequestBody struct {
	ID *string `json:"id"`
}

type ProductListRequestBody struct {
	ListFilter
	ProductRequestBody
	Order []OrderBy
}

type ProductListResponse = ListResponse[Product]

type ProductMsg struct {
	Single *request.RequestHandler[ProductRequestBody, Product]
	List   *request.RequestHandler[ProductListRequestBody, ProductListResponse]
	Create *request.RequestHandler[[]Product, request.Signal]
	Update *request.RequestHandler[Product, request.Signal]
	Delete *request.RequestHandler[string, request.Signal]
}
