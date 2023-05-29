package models

import (
	"time"

	"github.com/followthepattern/adapticc/pkg/request"
)

type LoginResponse struct {
	JWT       string    `json:"jwt,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"list_name"`
	Password  string `json:"password"`
}

type RegisterResponse struct {
	Email     *string `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

type AuthMsg struct {
	Register *request.RequestHandler[RegisterRequest, RegisterResponse]
	Login    *request.RequestHandler[LoginRequest, LoginResponse]
}
