package models

import (
	"github.com/followthepattern/adapticc/pkg/request"
)

type Product struct {
	ProductID   *string `db:"product_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (m *Product) IsNil() bool {
	return m == nil || m.ProductID == nil
}

type ProductRequestBody struct {
	ProductID *string `json:"product_id"`
}

type ProductListRequestBody struct {
	ListFilter
	ProductRequestBody
}

type ProductListResponse = ListResponse[Product]

type ProductMsg struct {
	Single *request.RequestHandler[ProductRequestBody, Product]
	List   *request.RequestHandler[ProductListRequestBody, ProductListResponse]
	Create *request.RequestHandler[[]Product, request.Signal]
	Update *request.RequestHandler[Product, request.Signal]
	Delete *request.RequestHandler[string, request.Signal]
}
