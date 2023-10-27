package models

import validation "github.com/go-ozzo/ozzo-validation"

type UserRole struct {
	UserID string `db:"user_id"`
	RoleID string `db:"role_id"`
	Userlog
}

func (m UserRole) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.UserID, validation.Required),
		validation.Field(&m.RoleID, validation.Required),
	)
}
