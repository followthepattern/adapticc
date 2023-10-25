package models

import validation "github.com/go-ozzo/ozzo-validation"

type Role struct {
	ID   string `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
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

func (r *Role) IsNil() bool {
	if r == nil {
		return true
	}

	if len(r.ID) < 1 {
		return true
	}

	return false
}

type RoleListRequestParams = ListRequestParams[ListFilter]

type RoleListResponse = ListResponse[Role]
