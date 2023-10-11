package utils

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"errors"

	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/google/uuid"
)

type ContextKey struct {
	Name string
}

var CtxUserKey = ContextKey{Name: "ctx-user"}

func GenerateHash(bytes []byte) []byte {
	hash := sha512.New512_256()
	hash.Write(bytes)
	return hash.Sum(nil)
}

func GenerateSalt() []byte {
	return GenerateHash([]byte(uuid.New().String()))
}

func GenerateSaltString() string {
	return hex.EncodeToString(GenerateSalt())
}

func GeneratePasswordHash(password string, salt string) string {
	return hex.EncodeToString(GenerateHash([]byte(password + salt)))
}

func GetModelFromContext[T any](ctx context.Context, ctxKey ContextKey) *T {
	obj := ctx.Value(ctxKey)

	model, ok := obj.(T)
	if !ok {
		return nil
	}

	return &model
}

func GetUserContext(ctx context.Context) (models.User, error) {
	obj := ctx.Value(CtxUserKey)

	model, ok := obj.(models.User)
	if !ok {
		return models.User{}, errors.New("invalid user context")
	}

	return model, nil
}
