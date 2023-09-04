package models

import (
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
