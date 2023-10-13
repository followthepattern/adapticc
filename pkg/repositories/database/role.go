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

func (repo Role) GetByID(id string) (*models.Role, error) {
	var data models.Role

	_, err := repo.db.From(S("usr").Table("roles")).
		Where(Ex{"id": id}).
		ScanStruct(&data)

	if err != nil {
		return nil, err
	}

	return &data, nil
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

func (repo Role) GetProfileRolesArray(userID string) ([]string, error) {
	roles, err := repo.GetRolesByUserID(userID)
	if err != nil {
		return nil, err
	}

	result := make([]string, len(roles))

	for i, role := range roles {
		result[i] = role.Name
	}

	return result, nil
}
