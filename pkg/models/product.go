package models

import (
	"github.com/followthepattern/adapticc/pkg/request"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	ID          *string `db:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Userlog
}

func (m *Product) IsNil() bool {
	return m == nil || m.ID == nil
}

func (e Product) CreateValidate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Title, validation.Required),
		validation.Field(&e.Description, validation.Required),
	)
}

func (e Product) UpdateValidate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.ID, validation.Required),
	)
}

type ProductListRequestParams = ListRequestParams[ListFilter]

type ProductListResponse = ListResponse[Product]

type ProductMsg struct {
	Single *request.Task[string, Product]
	List   *request.Task[ProductListRequestParams, ProductListResponse]
	Create *request.Task[[]Product, request.Signal]
	Update *request.Task[Product, request.Signal]
	Delete *request.Task[string, request.Signal]
}
