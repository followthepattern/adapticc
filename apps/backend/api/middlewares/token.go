package middlewares

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/types"
	"github.com/golang-jwt/jwt/v4"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer"
	InvalidToken        = "invalid token"
	ExpiredToken        = "ExpiredToken"
)

func invalidTokenError(field string) error {
	return fmt.Errorf("%s - missing field: %s", InvalidToken, field)
}

func getAuthorizedUserFromClaims(claims jwt.MapClaims) (*auth.AuthUser, error) {
	id, ok := claims["ID"].(string)
	if !ok {
		return nil, invalidTokenError("ID")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, invalidTokenError("email")
	}
	firstName, ok := claims["firstName"].(string)
	if !ok {
		return nil, invalidTokenError("firstName")
	}
	lastName, ok := claims["lastName"].(string)
	if !ok {
		return nil, invalidTokenError("lastName")
	}
	expiresAtStr, ok := claims["expiresAt"].(string)
	if !ok {
		return nil, invalidTokenError("expiresAt")
	}
	expiresAt, err := time.Parse(time.RFC3339, expiresAtStr)
	if err != nil {
		return nil, err
	}

	if time.Now().After(expiresAt) {
		return nil, errors.New(ExpiredToken)
	}

	return &auth.AuthUser{
		ID:        types.StringFrom(id),
		Email:     types.StringFrom(email),
		FirstName: types.StringFrom(firstName),
		LastName:  types.StringFrom(lastName),
	}, nil
}

func getToken(tokenString string) string {
	tokens := strings.Split(tokenString, BearerPrefix)

	if len(tokens) < 2 {
		return ""
	}

	return strings.TrimSpace(tokens[1])
}
