package database

import (
	"database/sql"
	"time"

	"log/slog"

	"github.com/followthepattern/adapticc/pkg/models"
	. "github.com/followthepattern/goqu/v9"
)

type Auth struct {
	db *Database
}

func NewAuth(database *sql.DB, logger *slog.Logger) Auth {
	db := New("postgres", database)

	return Auth{
		db: db,
	}
}

func (service Auth) VerifyEmail(email string) (bool, error) {
	count, err := service.db.From(userTableName).Where(Ex{"email": email}).Count()

	return count == 0, err
}

func (service Auth) RegisterUser(registerUser models.AuthUser) error {
	registerUser.Userlog = models.Userlog{
		CreatedAt: time.Now(),
	}

	_, err := service.db.Insert(userTableName).Rows(registerUser).Executor().Exec()
	return err
}

func (service Auth) VerifyLogin(email string) (models.AuthUser, error) {
	authUser := models.AuthUser{}

	_, err := service.db.From(userTableName).Where(Ex{"email": email}).ScanStruct(&authUser)
	if err != nil {
		return authUser, err
	}

	return authUser, err
}
