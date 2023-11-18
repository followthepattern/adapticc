package container

import (
	"database/sql"
	"log/slog"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/repositories/email"
)

type Container struct {
	ac      accesscontrol.AccessControl
	email   email.Email
	db      *sql.DB
	cfg     config.Config
	jwtKeys config.JwtKeyPair
	logger  *slog.Logger
}

func New(ac accesscontrol.AccessControl, emailClient email.Email, db *sql.DB, cfg config.Config, logger *slog.Logger, jwtKeys config.JwtKeyPair) Container {
	return Container{
		ac:      ac,
		db:      db,
		cfg:     cfg,
		logger:  logger,
		email:   emailClient,
		jwtKeys: jwtKeys,
	}
}

func (c Container) GetAccessControl() accesscontrol.AccessControl {
	return c.ac
}

func (c Container) GetDB() *sql.DB {
	return c.db
}

func (c Container) GetConfig() config.Config {
	return c.cfg
}

func (c Container) GetLogger() *slog.Logger {
	return c.logger
}

func (c Container) GetEmail() email.Email {
	return c.email
}

func (c Container) GetJWTKeys() config.JwtKeyPair {
	return c.jwtKeys
}
