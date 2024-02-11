package product

import (
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ProductModel struct {
	ID          types.String `db:"id" goqu:"omitempty"`
	Title       types.String `db:"title" goqu:"omitempty"`
	Description types.String `db:"description" goqu:"omitempty"`
	models.Userlog
}

func (m ProductModel) IsDefault() bool {
	return m.ID.Len() < 1
}

func (m ProductModel) CreateValidate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required),
		validation.Field(&m.Description, validation.Required),
	)
}

func (m ProductModel) UpdateValidate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ID, validation.Required),
	)
}

type ProductListRequestParams = models.ListRequestParams[models.ListFilter]

type ProductListResponse = models.ListResponse[ProductModel]
