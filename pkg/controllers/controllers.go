package controllers

import (
	"context"
	"database/sql"

	cerbos "github.com/cerbos/cerbos/client"
	"github.com/followthepattern/adapticc/pkg/config"
	"go.uber.org/zap"
)

type Controllers struct {
	user    User
	auth    Auth
	product Product
}

func New(ctx context.Context, cerbosClient cerbos.Client, db *sql.DB, cfg config.Config, logger *zap.Logger) Controllers {
	return Controllers{
		user:    NewUser(ctx, db, cfg, logger),
		auth:    NewAuth(ctx, db, cfg, logger),
		product: NewProduct(ctx, cerbosClient, db, cfg, logger),
	}
}

func (c Controllers) User() User {
	return c.user
}

func (c Controllers) Auth() Auth {
	return c.auth
}

func (c Controllers) Product() Product {
	return c.product
}
