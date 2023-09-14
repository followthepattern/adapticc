package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
	"go.uber.org/zap"

	. "github.com/doug-martin/goqu/v9"
)

type Auth struct {
	db  *Database
	ctx context.Context
}

func NewAuth(ctx context.Context, database *sql.DB, logger *zap.Logger) Auth {
	db := New("postgres", database)

	return Auth{
		ctx: ctx,
		db:  db,
	}
}

func (service Auth) VerifyEmail(email string) (bool, error) {
	count, err := service.db.From("usr.users").Where(Ex{"email": email}).Count()

	return count == 0, err
}

func (service Auth) RegisterUser(registerUser models.AuthUser) error {
	registerUser.Userlog = models.Userlog{
		CreatedAt: pointers.ToPtr(time.Now()),
	}

	_, err := service.db.Insert("usr.users").Rows(registerUser).Executor().Exec()
	return err
}

func (service Auth) VerifyLogin(email string) (models.AuthUser, error) {
	authUser := models.AuthUser{}

	_, err := service.db.From("usr.users").Where(Ex{"email": email}).ScanStruct(&authUser)

	return authUser, err
}
