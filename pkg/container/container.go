package container

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"

	"go.uber.org/zap"
)

type DependencyKey string

type ConstructorFunc[T any] func(cont *Container) (*T, error)

func getKey[T any](value T) DependencyKey {
	ttype := pointers.GetUnderlyingTypeRecursively(reflect.TypeOf(value))
	dependencyKey := DependencyKey(fmt.Sprintf("%s.%s", ttype.PkgPath(), ttype.Name()))
	return dependencyKey
}

func New(
	ctx context.Context,
	cfg config.Config,
	db *sql.DB,
	logger *zap.Logger) *Container {
	return &Container{
		dependencies: make(map[DependencyKey]any),
		ctx:          ctx,
		cfg:          cfg,
		db:           db,
		logger:       logger,
	}
}

type Container struct {
	ctx          context.Context
	cfg          config.Config
	db           *sql.DB
	dependencies map[DependencyKey]any
	logger       *zap.Logger
}

func Register[T any](cont *Container, constructor ConstructorFunc[T]) error {
	pointer, err := constructor(cont)
	if err != nil {
		return err
	}
	dependencyKey := getKey(pointer)
	cont.dependencies[dependencyKey] = pointer
	return nil
}

func RegisterWithCustomKey[T any](cont *Container, key DependencyKey, constructor ConstructorFunc[T]) error {
	pointer, err := constructor(cont)
	if err != nil {
		return err
	}
	cont.dependencies[key] = pointer
	return nil
}

func ResolveWithCustomKey[T any](cont *Container, key DependencyKey) (*T, error) {
	var dependency any
	dependency, ok := cont.dependencies[key]
	if !ok {
		return nil, fmt.Errorf("there is no registered object for this key: %v", key)
	}

	if result, ok := dependency.(*T); ok {
		return result, nil
	}
	return nil, fmt.Errorf("can't resolve %T", dependency)

}

func Resolve[T any](cont *Container) (*T, error) {
	var value T
	result, err := ResolveWithCustomKey[T](cont, getKey(value))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c Container) GetContext() context.Context {
	return c.ctx
}

func (c Container) GetConfig() config.Config {
	return c.cfg
}

func (c Container) GetDB() *sql.DB {
	return c.db
}

func (c Container) GetLogger() *zap.Logger {
	return c.logger
}
