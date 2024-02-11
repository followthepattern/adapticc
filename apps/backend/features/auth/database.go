package auth

import (
	"database/sql"

	"log/slog"

	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
	. "github.com/followthepattern/goqu/v9"
	"github.com/followthepattern/goqu/v9/exp"
)

var (
	userTableName = S("usr").Table("users")
)

type AuthDatabase struct {
	db *Database
}

func NewAuthDatabase(database *sql.DB, logger *slog.Logger) AuthDatabase {
	db := New("postgres", database)

	return AuthDatabase{
		db: db,
	}
}

func (service AuthDatabase) VerifyEmail(email types.String) (bool, error) {
	count, err := service.db.From(userTableName).Where(Ex{"email": email}).Count()

	return count == 0, err
}

func (service AuthDatabase) RegisterUser(registerUser AuthUser) error {
	registerUser.Userlog = models.Userlog{
		CreatedAt: types.TimeNow(),
	}

	_, err := service.db.Insert(userTableName).Rows(registerUser).Executor().Exec()
	return err
}

func (service AuthDatabase) VerifyLogin(email types.String) (AuthUser, error) {
	authUser := AuthUser{}

	_, err := service.db.From(userTableName).Where(Ex{"email": email}).ScanStruct(&authUser)
	if err != nil {
		return authUser, err
	}

	return authUser, err
}

func (repo AuthDatabase) GetRoleIDs(userID string) ([]string, error) {
	var roleIDs []string

	err := repo.db.From(S("usr").Table("user_role").As("ur")).
		Join(S("usr").Table("roles").As("r"),
			On(Ex{"r.id": I("ur.role_id")})).
		Where(Ex{"user_id": userID}).
		Select(T("r").Col(exp.Star())).
		ScanVals(roleIDs)

	if err != nil {
		return nil, err
	}

	return roleIDs, nil
}
