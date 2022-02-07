package container

import (
	"backend/internal/config"
	"backend/internal/utils"
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"go.uber.org/zap"
)

type IContainer interface {
	Register(key string, constructor func(IContainer) (interface{}, error)) error
	Resolve(string) (interface{}, error)
	Resolve2(target interface{}) error
	GetContext() *context.Context
	GetConfig() *config.Cfg
	GetDB() *sql.DB
	GetLogger() *zap.Logger
}

type Container struct {
	ctx          *context.Context
	cfg          *config.Cfg
	db           *sql.DB
	dependencies map[string]interface{}
	logger       *zap.Logger
}

func New(
	ctx *context.Context,
	cfg *config.Cfg,
	db *sql.DB,
	logger *zap.Logger) IContainer {
	return &Container{
		dependencies: make(map[string]interface{}),
		ctx:          ctx,
		cfg:          cfg,
		db:           db,
		logger:       logger,
	}
}

func (c *Container) Register(key string, constructor func(container IContainer) (interface{}, error)) error {
	obj, err := constructor(c)
	if err != nil {
		return err
	}
	c.dependencies[key] = obj
	return nil
}

func (c Container) Resolve(key string) (interface{}, error) {
	if dependeny, ok := c.dependencies[key]; ok {
		return dependeny, nil
	}
	return nil, fmt.Errorf("there is no registered object for this key: %v", key)
}

func (c Container) Resolve2(target interface{}) error {
	rv := reflect.ValueOf(target)

	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("target is not a pointer")
	}
	key := utils.GetKey(target)
	dependency, ok := c.dependencies[key]
	if !ok {
		return fmt.Errorf("there is no registered object for this key: %v", key)
	}
	target = &dependency

	return nil
}

func (c Container) GetContext() *context.Context {
	return c.ctx
}

func (c Container) GetConfig() *config.Cfg {
	return c.cfg
}

func (c Container) GetDB() *sql.DB {
	return c.db
}

func (c Container) GetLogger() *zap.Logger {
	return c.logger
}
