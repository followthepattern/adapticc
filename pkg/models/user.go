package models

import (
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
)

type User struct {
	ID        *string `json:"id,omitempty" goqu:"skipupdate"`
	Email     *string `json:"email,omitempty" goqu:"skipupdate"`
	FirstName *string `json:"first_name,omitempty" db:"first_name"`
	LastName  *string `json:"last_name,omitempty" db:"last_name"`
	Active    *bool   `json:"active,omitempty" goqu:"skipupdate"`
	Userlog
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
