package models

import (
	"time"

	"github.com/followthepattern/adapticc/types"
	"github.com/followthepattern/adapticc/user"
	validation "github.com/go-ozzo/ozzo-validation"
)

type LoginResponse struct {
	JWT       string    `json:"jwt,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}

type LoginRequestParams struct {
	Email    types.String `json:"email"`
	Password types.String `json:"password"`
}

func (l LoginRequestParams) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required),
		validation.Field(&l.Password, validation.Required),
	)
}

type RegisterRequestParams struct {
	Email     types.String `json:"email"`
	FirstName types.String `json:"first_name"`
	LastName  types.String `json:"list_name"`
	Password  types.String `json:"password"`
}

func (r RegisterRequestParams) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.FirstName, validation.Required),
		validation.Field(&r.LastName, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

type Password struct {
	PasswordHash string `db:"password_hash"`
}

func (p *Password) IsEmpty() bool {
	if p == nil {
		return true
	}

	if len(p.PasswordHash) < 1 {
		return true
	}

	return false
}

type AuthUser struct {
	user.UserModel
	Password
}

type RegisterResponse struct {
	Email     types.String `json:"email"`
	FirstName types.String `json:"first_name"`
	LastName  types.String `json:"last_name"`
}
