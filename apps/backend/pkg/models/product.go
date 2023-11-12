package models

import (
	"github.com/followthepattern/adapticc/pkg/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	ID          types.NullString `db:"id" goqu:"omitempty"`
	Title       types.NullString `db:"title" goqu:"omitempty"`
	Description types.NullString `db:"description" goqu:"omitempty"`
	Userlog
}

func (m Product) IsDefault() bool {
	return m.ID.Len() < 1
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
