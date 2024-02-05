package models

import (
	"github.com/followthepattern/adapticc/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Role struct {
	ID   types.String `db:"id"`
	Code types.String `db:"code"`
	Name types.String `db:"name"`
	Userlog
}

func (m Role) CreateValidate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Code, validation.Required),
		validation.Field(&m.Name, validation.Required),
	)
}

func (m Role) UpdateValidate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ID, validation.Required),
	)
}

func (m Role) IsDefault() bool {
	return m.ID.Len() < 1
}

type RoleListRequestParams = ListRequestParams[ListFilter]

type RoleListResponse = ListResponse[Role]
