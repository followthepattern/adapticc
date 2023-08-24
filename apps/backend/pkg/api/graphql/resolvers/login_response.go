package resolvers

import (
	"github.com/followthepattern/adapticc/pkg/models"

	"github.com/followthepattern/graphql-go"
)

type loginResponse struct {
	JWT       string       `json:"jwt,omitempty"`
	ExpiresAt graphql.Time `json:"expires_at,omitempty"`
}

func getFromLoginResponseModel(m models.LoginResponse) *loginResponse {
	expiresAt := graphql.Time{
		Time: m.ExpiresAt,
	}
	return &loginResponse{
		ExpiresAt: expiresAt,
		JWT:       m.JWT,
	}
}
