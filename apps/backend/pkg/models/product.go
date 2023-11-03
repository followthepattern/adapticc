package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	ID          string `db:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Userlog
}

func (m Product) IsDefault() bool {
	return len(m.ID) < 1
}

func (m Product) CreateValidate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required),
		validation.Field(&m.Description, validation.Required),
	)
}

func (m Product) UpdateValidate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ID, validation.Required),
	)
}

type ProductListRequestParams = ListRequestParams[ListFilter]

type ProductListResponse = ListResponse[Product]
