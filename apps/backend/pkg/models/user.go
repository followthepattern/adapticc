package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	ID        string `db:"id" goqu:"skipupdate,omitempty"`
	Email     string `db:"email" goqu:"skipupdate,omitempty"`
	FirstName string `db:"first_name" goqu:"omitempty"`
	LastName  string `db:"last_name" goqu:"omitempty"`
	Active    bool   `db:"active" goqu:"skipupdate,omitempty"`
	Roles     []Role `db:"-"`
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
	return len(u.ID) < 1
}

type SingleUserRequestParams struct {
	ID    *string
	Email *string
}

type UserListRequestParams = ListRequestParams[ListFilter]

type UserListResponse = ListResponse[User]
