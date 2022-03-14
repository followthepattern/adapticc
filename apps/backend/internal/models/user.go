package models

import (
	"backend/internal/utils/pointers"
	"time"
)

type User struct {
	Userlog
	ID           *string    `json:"id,omitempty"`
	Email        *string    `json:"email,omitempty"`
	FirstName    *string    `json:"first_name,omitempty" db:"first_name"`
	LastName     *string    `json:"last_name,omitempty" db:"last_name"`
	PasswordHash *string    `json:"password_hash,omitempty" db:"password_hash"`
	Salt         *string    `json:"salt,omitempty"`
	Active       *bool      `json:"active,omitempty"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
}

func (u User) IsNil() bool {
	return u.ID == nil
}

var AnnonymusUser User = User{
	ID: pointers.String(""),
}
