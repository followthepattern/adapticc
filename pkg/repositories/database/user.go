package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database/sqlbuilder"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"

	. "github.com/doug-martin/goqu/v9"
)

type User struct {
	db  *Database
	ctx context.Context
}

func (User) tableName() string {
	return "usr.users"
}

func NewUser(ctx context.Context, database *sql.DB) User {
	db := New("postgres", database)

	return User{
		ctx: ctx,
		db:  db,
	}
}

func (repo User) Create(userID string, users []models.User) (err error) {
	count, err := sqlbuilder.GetInsertWithPermissions(repo.db, "USER", userID)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("there is no effective permission to create this resource")
	}

	active := false

	for i := range users {
		users[i].Userlog = setCreateUserlog(userID, time.Now())
		users[i].Active = &active
	}
	_, err = repo.db.Insert(repo.tableName()).Rows(users).Executor().Exec()
	return
}

func (repo User) GetByID(id string) (*models.User, error) {
	user := models.User{}

	query := repo.db.From(repo.tableName()).Where(Ex{"id": id})

	_, err := query.ScanStruct(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (repo User) GetByEmail(email string) (*models.User, error) {
	user := models.User{}

	query := repo.db.From(repo.tableName()).Where(Ex{"email": email})

	_, err := query.ScanStruct(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (repo User) Get(userID string, request models.UserListRequestParams) (*models.UserListResponse, error) {
	data := []models.User{}

	query := repo.db.From(repo.tableName())

	if request.Filter.Search != nil {
		pattern := fmt.Sprintf("%%%s%%", *request.Filter.Search)
		query = query.Where(
			Or(
				I("first_name").Like(pattern),
				I("last_name").Like(pattern),
				I("email").Like(pattern),
			))
	}

	query = sqlbuilder.GetSelectWithPermissions(
		query,
		"USER",
		I("users.id"),
		userID,
	)

	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	if request.Pagination.Page == nil {
		request.Pagination.Page = pointers.ToPtr[uint](models.DefaultPage)
	}

	if request.Pagination.PageSize != nil {
		page := *request.Pagination.Page
		if page > 0 {
			page--
		}

		query = query.Offset(page * *request.Pagination.PageSize)
		query = query.Limit(*request.Pagination.PageSize)
	}

	err = query.ScanStructs(&data)
	if err != nil {
		return nil, err
	}

	result := models.UserListResponse{
		Count:    count,
		PageSize: request.Pagination.PageSize,
		Page:     request.Pagination.Page,
		Data:     data,
	}

	return &result, nil
}

func (repo User) Update(userID string, user models.User) error {
	user.Userlog = setUpdateUserlog(userID, time.Now())

	query := repo.db.Update(repo.tableName()).
		Set(user).
		Where(C("id").Eq(*user.ID))

	query = sqlbuilder.GetUpdateWithPermissions(
		query,
		"USER",
		I("users.id"),
		userID,
	)

	_, err := query.Executor().Exec()
	return err
}

func (repo User) Delete(userID, id string) error {
	query := repo.db.Delete(repo.tableName()).
		Where(C("id").Eq(id))

	query = sqlbuilder.GetDeleteWithPermissions(
		query,
		"USER",
		I("usr.users.id"),
		userID,
	)

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

	return err
}
