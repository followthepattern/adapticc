package models

import (
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	ID        *string `db:"id" goqu:"skipupdate"`
	Email     *string `db:"email" goqu:"skipupdate"`
	FirstName *string `db:"first_name"`
	LastName  *string `db:"last_name"`
	Active    *bool   `db:"active" goqu:"skipupdate"`
	Roles     []Role  `db:"-"`
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

func (u User) IsNil() bool {
	return u.ID == nil
}

func (u User) IsDefault() bool {
	if u.IsNil() {
		return true
	}

	return len(*u.ID) < 1
}

var UnAuthorizedUser User = User{
	ID: pointers.ToPtr(""),
}

type SingleUserRequestParams struct {
	ID    *string
	Email *string
}

type UserListRequestParams = ListRequestParams[ListFilter]

type UserListResponse = ListResponse[User]
