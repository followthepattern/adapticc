package auth

import (
	"time"

	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
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

type AuthUser struct {
	ID           types.String `db:"id" goqu:"skipupdate,omitempty"`
	Email        types.String `db:"email" goqu:"skipupdate,omitempty"`
	FirstName    types.String `db:"first_name" goqu:"omitempty"`
	LastName     types.String `db:"last_name" goqu:"omitempty"`
	Active       types.Bool   `db:"active" goqu:"skipupdate,omitempty"`
	PasswordHash string       `db:"password_hash"`
	models.Userlog
}

func (model AuthUser) IsDefault() bool {
	return !model.Email.IsValid() || !model.FirstName.IsValid() || !model.LastName.IsValid()
}

type RegisterResponse struct {
	Email     types.String `json:"email"`
	FirstName types.String `json:"first_name"`
	LastName  types.String `json:"last_name"`
}
