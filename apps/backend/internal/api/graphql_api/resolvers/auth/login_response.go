package auth

import (
	"backend/internal/models"

	"github.com/graph-gophers/graphql-go"
)

type loginResponse struct {
	JWT       *string       `json:"jwt,omitempty"`
	ExpiresAt *graphql.Time `json:"expires_at,omitempty"`
}

func getFromLoginResponseModel(m models.LoginResponse) *loginResponse {
	expiresAt := graphql.Time{
		Time: *m.ExpiresAt,
	}
	return &loginResponse{
		JWT:       m.JWT,
		ExpiresAt: &expiresAt,
	}
}
