package auth

import (
	"context"
	"errors"
)

type ContextKey struct {
	Name string
}

var CtxUserKey = ContextKey{Name: "ctx-user"}

func GetModelFromContext[T any](ctx context.Context, ctxKey ContextKey) *T {
	obj := ctx.Value(ctxKey)

	model, ok := obj.(T)
	if !ok {
		return nil
	}

	return &model
}

func GetUserContext(ctx context.Context) (AuthUser, error) {
	obj := ctx.Value(CtxUserKey)

	model, ok := obj.(AuthUser)
	if !ok {
		return AuthUser{}, errors.New("invalid user context")
	}

	if model.IsDefault() {
		return AuthUser{}, errors.New("invalid user context")
	}

	return model, nil
}
