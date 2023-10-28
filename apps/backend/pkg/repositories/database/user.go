package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"

	. "github.com/doug-martin/goqu/v9"
)

type User struct {
	db  *Database
	ctx context.Context
}

var (
	userTableName = S("usr").Table("users")
)

func NewUser(ctx context.Context, database *sql.DB) User {
	db := New("postgres", database)

	return User{
		ctx: ctx,
		db:  db,
	}
}

func (repo User) Create(users []models.User) (err error) {
	for i := range users {
		users[i].Userlog.CreatedAt = pointers.ToPtr(time.Now())
	}
	_, err = repo.db.Insert(userTableName).Rows(users).Executor().Exec()
	return
}

func (repo User) GetByID(id string) (*models.User, error) {
	user := models.User{}

	query := repo.db.From(userTableName).Where(Ex{"id": id})

	_, err := query.ScanStruct(&user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (repo User) GetByEmail(email string) (*models.User, error) {
	user := models.User{}

	query := repo.db.From(userTableName).Where(Ex{"email": email})

	_, err := query.ScanStruct(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (repo User) Get(request models.UserListRequestParams) (*models.UserListResponse, error) {
	data := []models.User{}

	query := repo.db.From(userTableName)

	if request.Filter.Search != nil {
		pattern := fmt.Sprintf("%%%s%%", *request.Filter.Search)
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

func (repo User) Update(user models.User) error {
	user.Userlog.UpdatedAt = pointers.ToPtr(time.Now())

	query := repo.db.Update(userTableName).
		Set(user).
		Where(C("id").Eq(*user.ID))

	_, err := query.Executor().Exec()
	return err
}

func (repo User) ActivateUser(userID string) error {
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

func (repo User) Delete(id string) error {
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
