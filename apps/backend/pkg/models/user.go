package models

import (
	"time"

	"github.com/followthepattern/adapticc/pkg/request"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
)

type UserRequestBody struct {
	ID    *string
	Email *string
}

type UserListRequestBody struct {
	ListFilter
	UserRequestBody
}

type User struct {
	ID           *string    `json:"id,omitempty" goqu:"skipupdate"`
	Email        *string    `json:"email,omitempty" goqu:"skipupdate"`
	FirstName    *string    `json:"first_name,omitempty" db:"first_name"`
	LastName     *string    `json:"last_name,omitempty" db:"last_name"`
	Password     *string    `db:"password" goqu:"skipupdate"`
	Salt         *string    `json:"salt,omitempty" goqu:"skipupdate"`
	Active       *bool      `json:"active,omitempty" goqu:"skipupdate"`
	RegisteredAt *time.Time `json:"registered_at,omitempty" db:"registered_at" goqu:"skipupdate"`
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

var Guest User = User{
	ID: pointers.String(""),
}

type UserListResponse = ListResponse[User]

type UserMsg struct {
	Single *request.RequestHandler[UserRequestBody, User]
	List   *request.RequestHandler[UserListRequestBody, UserListResponse]
	Create *request.RequestHandler[[]User, request.Signal]
	Update *request.RequestHandler[User, request.Signal]
	Delete *request.RequestHandler[string, request.Signal]
}
