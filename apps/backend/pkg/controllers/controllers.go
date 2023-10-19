package controllers

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/config"
)

type Controllers struct {
	user    User
	auth    Auth
	product Product
}

func New(ctx context.Context, ac accesscontrol.AccessControl, db *sql.DB, cfg config.Config, logger *slog.Logger) Controllers {
	return Controllers{
		user:    NewUser(ctx, ac, db, cfg, logger),
		auth:    NewAuth(ctx, db, cfg, logger),
		product: NewProduct(ctx, ac, db, cfg, logger),
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
