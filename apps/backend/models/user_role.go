package models

import (
	"github.com/followthepattern/adapticc/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserRole struct {
	UserID types.String `db:"user_id"`
	RoleID types.String `db:"role_id"`
	Userlog
}

func (m UserRole) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.UserID, validation.Required),
		validation.Field(&m.RoleID, validation.Required),
	)
}
