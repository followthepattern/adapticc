package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/followthepattern/adapticc/repositories/database/sqlbuilder"
	"github.com/followthepattern/adapticc/types"

	. "github.com/followthepattern/goqu/v9"
)

type UserDatabase struct {
	db *Database
}

var (
	userTableName = S("usr").Table("users")
)

func NewUserDatabase(database *sql.DB) UserDatabase {
	db := New("postgres", database)

	return UserDatabase{
		db: db,
	}
}

func (repo UserDatabase) Create(users []UserModel) (err error) {
	for i := range users {
		users[i].Userlog.CreatedAt = types.TimeNow()
	}
	_, err = repo.db.Insert(userTableName).Rows(users).Executor().Exec()
	return
}

func (repo UserDatabase) GetByID(id types.String) (*UserModel, error) {
	user := UserModel{}

	query := repo.db.From(userTableName).Where(Ex{"id": id})

	_, err := query.ScanStruct(&user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (repo UserDatabase) GetByEmail(email string) (*UserModel, error) {
	user := UserModel{}

	query := repo.db.From(userTableName).Where(Ex{"email": email})

	_, err := query.ScanStruct(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (repo UserDatabase) Get(request UserListRequestParams) (*UserListResponse, error) {
	data := []UserModel{}

	query := repo.db.From(userTableName)

	if request.Filter.Search.IsValid() {
		pattern := fmt.Sprintf("%%%s%%", request.Filter.Search)
		query = query.Where(
			Or(
				I("first_name").Like(pattern),
				I("last_name").Like(pattern),
				I("email").Like(pattern),
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

	result := UserListResponse{
		Count:    types.Int64From(count),
		PageSize: request.Pagination.PageSize,
		Page:     request.Pagination.Page,
		Data:     data,
	}

	return &result, nil
}

func (repo UserDatabase) Update(user UserModel) error {
	user.Userlog.UpdatedAt = types.TimeNow()

	query := repo.db.Update(userTableName).
		Set(user).
		Where(C("id").Eq(user.ID))

	_, err := query.Executor().Exec()
	return err
}

func (repo UserDatabase) ActivateUser(userID string) error {
	result, err := repo.db.Update(userTableName).
		Set(Record{"active": true}).
		Where(Ex{"id": userID, "active": false}).
		Executor().Exec()
	if err != nil {
		return err
	}

	effectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if effectedRows == 0 {
		return fmt.Errorf("there is no existing inactive user with ID: %s", userID)
	}

	return nil
}

func (repo UserDatabase) Delete(id types.String) error {
	query := repo.db.Delete(userTableName).
		Where(C("id").Eq(id))

	result, err := query.Executor().Exec()
	if err != nil {
		return err
	}

	effectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if effectedRows == 0 {
		return errors.New("no rows have been effected")
	}

	return nil
}
