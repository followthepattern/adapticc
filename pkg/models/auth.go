package models

import (
	"time"
)

type LoginResponse struct {
	JWT       string    `json:"jwt,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}

type LoginRequestParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequestParams struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"list_name"`
	Password  string `json:"password"`
}

type Password struct {
	PasswordHash string `db:"password_hash"`
	Salt         string `db:"salt"`
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
	User
	Password
}

type RegisterResponse struct {
	Email     *string `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}