package database

import (
	"context"
	"database/sql"

	. "github.com/doug-martin/goqu/v9"
	"github.com/followthepattern/adapticc/pkg/models"
)

type Role struct {
	db  *Database
	ctx context.Context
}

func NewRole(ctx context.Context, database *sql.DB) Role {
	db := New("postgres", database)

	return Role{
		ctx: ctx,
		db:  db,
	}
}

func (repo Role) GetRolesByUserID(userID string) ([]models.Role, error) {
	var data []models.Role

	err := repo.db.From(S("usr").Table("user_role").As("ur")).
		Join(S("usr").Table("roles").As("r"),
			On(Ex{"r.id": I("ur.role_id")})).
		Where(Ex{"user_id": userID}).
		ScanStructs(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
