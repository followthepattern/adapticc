package middlewares

import (
	"context"
	"errors"
	"time"

	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer"
)

var CtxUserKey = &utils.ContextKey{Name: "ctx-user"}
var MachineUserKey = &utils.ContextKey{Name: "ctx-machine"}

func getUserContextFromClaims(claims jwt.MapClaims) (*models.User, error) {
	id, ok := claims["ID"].(string)
	if !ok {
		return nil, errors.New("ID is not in claims")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("email is not in claims")
	}
	firstName, ok := claims["firstName"].(string)
	if !ok {
		return nil, errors.New("firstName is not in claims")
	}
	lastName, ok := claims["lastName"].(string)
	if !ok {
		return nil, errors.New("lastName is not in claims")
	}
	expiresAtStr, ok := claims["expiresAt"].(string)
	if !ok {
		return nil, errors.New("expiresAt is not in claims")
	}

	expiresAt, err := time.Parse(time.RFC3339, expiresAtStr)
	if err != nil {
		return nil, errors.New("expiresAt doesn't have the right time format")
	}

	if time.Now().After(expiresAt) {
		return nil, errors.New("token is expired")
	}

	user := models.User{
		ID:        &id,
		Email:     &email,
		FirstName: &firstName,
		LastName:  &lastName,
	}
	return &user, nil
}

func GetUserFromContext(ctx context.Context) *models.User {
	obj := ctx.Value(CtxUserKey)

	user, ok := obj.(*models.User)
	if !ok {
		return nil
	}

	return user
}
