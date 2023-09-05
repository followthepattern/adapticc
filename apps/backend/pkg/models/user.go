package models

import (
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
)

type User struct {
	ID        *string `db:"id" goqu:"skipupdate"`
	Email     *string `db:"email" goqu:"skipupdate"`
	FirstName *string `db:"first_name"`
	LastName  *string `db:"last_name"`
	Active    *bool   `db:"active" goqu:"skipupdate"`
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
