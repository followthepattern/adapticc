package models

import (
	"github.com/followthepattern/adapticc/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	ID        types.String `db:"id" goqu:"skipupdate,omitempty"`
	Email     types.String `db:"email" goqu:"skipupdate,omitempty"`
	FirstName types.String `db:"first_name" goqu:"omitempty"`
	LastName  types.String `db:"last_name" goqu:"omitempty"`
	Active    types.Bool   `db:"active" goqu:"skipupdate,omitempty"`
	Roles     []Role       `db:"-"`
	Userlog
}

func (u User) CreateValidate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.FirstName, validation.Required),
		validation.Field(&u.LastName, validation.Required),
	)
}

func (u User) UpdateValidate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.ID, validation.Required),
	)
}

func (u User) IsDefault() bool {
	return u.ID.Len() < 1
}

type UserListRequestParams = ListRequestParams[ListFilter]

type UserListResponse = ListResponse[User]
