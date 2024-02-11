package role

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/followthepattern/adapticc/repositories/database/sqlbuilder"
	"github.com/followthepattern/adapticc/types"
	. "github.com/followthepattern/goqu/v9"
	"github.com/followthepattern/goqu/v9/exp"
)

var (
	roleTableName     = S("usr").Table("roles")
	userRoleTableName = S("usr").Table("user_role")
)

type RoleDatabase struct {
	db *Database
}

func NewRoleDatabase(database *sql.DB) RoleDatabase {
	db := New("postgres", database)

	return RoleDatabase{
		db: db,
	}
}

func (repo RoleDatabase) GetByID(id string) (*RoleModel, error) {
	var data RoleModel

	_, err := repo.db.From(roleTableName).
		Where(Ex{"id": id}).
		ScanStruct(&data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo RoleDatabase) Get(request RoleListRequestParams) (*RoleListResponse, error) {
	data := []RoleModel{}

	query := repo.db.From(roleTableName)

	if request.Filter.Search.IsValid() {
		pattern := fmt.Sprintf("%%%s%%", request.Filter.Search)
		query = query.Where(
			Or(
				I("id").Like(pattern),
				I("name").Like(pattern),
				I("code").Like(pattern),
			))
	}

	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	query = sqlbuilder.WithPagination(query, request.Pagination)

	query = sqlbuilder.WithOrderBy(query, request.OrderBy)

	err = query.ScanStructs(&data)
	if err != nil {
		return nil, err
	}

	result := RoleListResponse{
		Count:    types.Int64From(count),
		PageSize: request.Pagination.PageSize,
		Page:     request.Pagination.Page,
		Data:     data,
	}

	return &result, nil
}

func (repo RoleDatabase) GetRolesByUserID(userID types.String) ([]RoleModel, error) {
	var data []RoleModel

	err := repo.db.From(userRoleTableName.As("ur")).
		Join(roleTableName.As("r"),
			On(Ex{"r.id": I("ur.role_id")})).
		Where(Ex{"user_id": userID}).
		Select(T("r").Col(exp.Star())).
		ScanStructs(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo RoleDatabase) AddRoleToUser(values []UserRoleModel) error {
	for i, _ := range values {
		values[i].Userlog.CreatedAt = types.TimeNow()
	}

	insertion := repo.db.Insert(userRoleTableName)

	_, err := insertion.Rows(values).Executor().Exec()
	return err
}

func (repo RoleDatabase) RemoveRoleFromUser(value UserRoleModel) error {
	res, err := repo.db.
		Delete(userRoleTableName).
		Where(Ex{
			"user_id": value.UserID,
			"role_id": value.RoleID}).
		Executor().
		Exec()

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows < 1 {
		return errors.New("no rows been deleted")
	}

	return err
}
